package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	"github.com/lios/go-activiti/logger"
)

var (
	variableEntity      VariableEntity
	variableDataManager VariableDataManager
)

type VariableEntityManagerImpl struct {
	AbstractEntityManager
}

//func init() {
//	variableDataManager = VariableDataManager{AbstractDataManager:AbstractDataManager{TableModel{AbstractModel(Variable{})}}}
//}
func (variableEntityManager VariableEntityManagerImpl) GetDataManager() DataManagers {
	return variableDataManager
}

func (variableEntityManager VariableEntityManagerImpl) DeleteVariableInstanceByTask(taskId int64) {
	dataManager := variableEntityManager.GetDataManager()
	variableDataManager := dataManager.(VariableDataManager)
	variables, err := variableDataManager.SelectByTaskId(taskId)
	if err != nil {
		logger.Error("select by taskId err：", err)
		return
	}
	for _, variable := range variables {
		variableDataManager.Delete(variable.Id)
	}
}
