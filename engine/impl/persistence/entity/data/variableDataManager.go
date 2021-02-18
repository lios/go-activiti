package data

import (
	"github.com/lios/go-activiti/db"
	. "github.com/lios/go-activiti/engine/variable"
	"github.com/lios/go-activiti/errs"
	"github.com/prometheus/common/log"
)

type VariableDataManager struct {
	Variable
	AbstractDataManager
}

func (define VariableDataManager) Create(name string, variableType VariableType, value interface{}) *Variable {
	variable := Variable{}
	variable.Version = 0
	variable.Name = name
	variable.Type = variableType.GetTypeName()
	variable.SetValue(value, variableType)
	return &variable

}

func (defineManager VariableDataManager) createHistoricVariable() (err error) {
	variable := defineManager.Variable
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

func (defineManager VariableDataManager) SelectProcessInstanceId(name string, processInstanceId int64) (Variable, error) {
	variables := Variable{}
	err := db.DB().Where("proc_inst_id = ?", processInstanceId).Where("name = ?", name).First(&variables).Error
	if err != nil {
		log.Infoln("Select Variable err: ", err)
		return Variable{}, err
	}
	return variables, nil
}

func (variableManager VariableDataManager) SelectTaskId(name string, taskId int64) (Variable, error) {
	variables := Variable{}
	err := db.DB().Where("task_id = ?", taskId).Where("name = ?", name).First(&variables).Error
	if err != nil {
		log.Infoln("根据[taskId] 查询流程变量异常", err)
		return Variable{}, err
	}
	return variables, nil
}

func (variableManager VariableDataManager) SelectByProcessInstanceId(processInstanceId int64) ([]Variable, error) {
	variables := make([]Variable, 0)
	err := db.DB().Where("proc_inst_id = ?", processInstanceId).Find(&variables).Error
	if err != nil {
		log.Infoln("Select Variable err: ", err)
		return variables, err
	}
	if variables != nil && len(variables) > 0 {
		return variables, nil
	}
	return variables, errs.ProcessError{Code: "1001", Msg: "Not Find"}
}

func (variableManager VariableDataManager) SelectByTaskId(taskId int64) ([]Variable, error) {
	variables := make([]Variable, 0)
	err := db.DB().Where("task_id = ?", taskId).Find(&variables).Error
	if err != nil {
		log.Infoln("Select Variable err: ", err)
		return variables, err
	}
	if variables != nil && len(variables) > 0 {
		return variables, nil
	}
	return variables, errs.ProcessError{Code: "1001", Msg: "Not Find"}
}

func (variableManager VariableDataManager) Delete(variableId int64) error {
	variable := Variable{}
	err := db.DB().Where("id=?", variableId).Delete(variable).Error
	if err != nil {
		log.Infoln("delete Variable err: ", err)
	}
	return err
}
