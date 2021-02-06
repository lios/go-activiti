package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	. "github.com/lios/go-activiti/engine/variable"
	"github.com/lios/go-activiti/errs"
	"reflect"
)

type AbstractVariableScopeImpl interface {
	GetSourceActivityExecution() ExecutionEntity
}
type VariableScopeImpl struct {
	AbstractVariableScopeImpl
}

func (variableScope VariableScopeImpl) SetVariableLocal(variables map[string]interface{}) error {
	variableManager := GetVariableManager()
	variableTypes := variableManager.VariableTypes
	manager := GetVariableEntityManager()
	dataManager := manager.GetDataManager()
	execution := variableScope.GetSourceActivityExecution()
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
			err := variableDataManager.Insert(&variable)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
