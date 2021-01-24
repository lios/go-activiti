package handler

import (
	. "github.com/lios/go-activiti/engine"
	"github.com/lios/go-activiti/errs"
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

type ActivitiConstructor func(entity ExecutionEntity) IActiviti

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
