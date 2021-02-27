package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	. "github.com/lios/go-activiti/engine/variable"
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
func (variableEntityManager VariableEntityManagerImpl) CreteVariable(variable Variable) error {
	dataManager := variableEntityManager.GetDataManager()
	dataManager.Insert(&variable)
	variableEntityManager.CreateHistoricVariable(variable)
	return nil
}

func (variableEntityManager VariableEntityManagerImpl) CreateHistoricVariable(variable Variable) (err error) {
	historicVariable := HistoricVariable{}

	historicVariable.TaskId = variable.TaskId
	historicVariable.ProcessInstanceId = variable.ProcessInstanceId
	historicVariable.Name = variable.Name
	historicVariable.Version = variable.Version
	historicVariable.Type = variable.Type
	historicVariable.Text = variable.Text
	historicVariable.Number = variable.Number
	historicVariable.Date = variable.Date
	historicVariable.Float = variable.Float
	historicVariable.Blob = variable.Blob

	historicVariableManager := HistoricVariableDataManager{}
	historicVariableManager.HistoricVariable = historicVariable
	return historicVariableManager.Insert()
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
