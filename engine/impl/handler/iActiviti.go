package handler

import (
	. "github.com/lios/go-activiti/engine/contanst"
	"github.com/lios/go-activiti/engine/delegate"
	"github.com/lios/go-activiti/errs"
	"reflect"
	"sync"
)

var gConstructorMap map[string]ActivitiConstructor
var lock sync.Mutex

func init() {
	gConstructorMap = make(map[string]ActivitiConstructor, 0)
}

type IActiviti interface {
	GetInPut() interface{}
	GetOutPut() interface{}
}

type ActivitiConstructor func(entity delegate.DelegateExecution) IActiviti

func RegisterConstructor(name string, constructor ActivitiConstructor) error {
	lock.Lock()
	defer lock.Unlock()
	_, ok := gConstructorMap[name]
	if !ok {
		gConstructorMap[name] = constructor
	} else {
		return errs.ProcessError{Code: "1005", Msg: "name has register"}
	}
	return nil
}

func GetConstructorByName(name string) (ActivitiConstructor, error) {
	lock.Lock()
	defer lock.Unlock()
	constructor, ok := gConstructorMap[name]
	if !ok {
		return nil, errs.ProcessError{Code: "1006", Msg: "name not find"}
	}
	return constructor, nil
}

func PerformTaskListener(entity delegate.DelegateExecution, taskName, processKey string) error {
	activitiConstructor, err := GetConstructorByName(processKey)
	if err != nil {
		return err
	}
	constructor := activitiConstructor(entity)
	reflectConstructor := reflect.ValueOf(constructor)
	taskParams := []reflect.Value{reflectConstructor}

	method, b := reflectConstructor.Type().MethodByName(taskName)
	if !b {
		return nil
	}

	callResponse := method.Func.Call(taskParams)
	code := callResponse[0].Interface()
	errRes := callResponse[1].Interface()
	code = code.(string)
	if code != ACTIVITI_HANDLER_CODE {
		err := errRes.(error)
		return err
	}
	return nil
}
