package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	. "github.com/lios/go-activiti/engine/variable"
	"github.com/lios/go-activiti/errs"
	"reflect"
	"sync"
)

var usedVariablesCache *sync.Pool

type AbstractVariableScopeImpl interface {
	GetSourceActivityExecution() ExecutionEntity
}

func init() {
	usedVariablesCache = &sync.Pool{
		New: func() interface{} {
			return nil
		},
	}
}

type VariableScopeImpl struct {
	AbstractVariableScopeImpl
	ExecutionEntity ExecutionEntity
}

//保存流程变量
func (variableScope VariableScopeImpl) SetVariable(execution ExecutionEntity, variables map[string]interface{}) error {
	variableManager := GetVariableManager()
	variableEntityManager := GetVariableEntityManager()
	variableTypes := variableManager.VariableTypes
	manager := GetVariableEntityManager()
	dataManager := manager.GetDataManager()
	variableDataManager := dataManager.(VariableDataManager)
	if variables != nil && len(variables) > 0 {
		for k, v := range variables {
			kind := reflect.TypeOf(v).Kind()
			variableType := variableTypes.GetVariableType(kind.String())
			if variableType == nil {
				return errs.ProcessError{Code: "1001", Msg: "no type"}
			}
			variable := variableDataManager.Create(k, variableType, v)
			//存在更新
			specificVariable, e := execution.GetSpecificVariable(k)
			if e != nil {
				variable.Version = specificVariable.Version + 1
			}

			execution.SetScope(variable)
			err := variableEntityManager.CreteVariable(*variable)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (variableScope VariableScopeImpl) SetVariableLocal(variables map[string]interface{}) error {
	variable := usedVariablesCache.Get()
	if variable == nil {
		usedVariablesCache.Put(variables)
	} else {
		variablesCache := variable.(map[string]interface{})
		for k, v := range variables {
			variablesCache[k] = v
		}
		usedVariablesCache.Put(variablesCache)
	}
	return nil
}

func (variableScope VariableScopeImpl) GetVariableLocal() (variables map[string]interface{}) {
	variable := usedVariablesCache.Get()
	if variable == nil {
		return nil
	} else {
		variablesCache := variable.(map[string]interface{})
		return variablesCache
	}
	return nil
}
