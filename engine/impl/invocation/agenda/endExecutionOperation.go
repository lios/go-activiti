package agenda

import (
	"github.com/lios/go-activiti/engine/impl/manager"
)

type EndExecutionOperation struct {
	AbstractOperation
}

func (end *EndExecutionOperation) Run() (err error) {
	err = deleteDataForExecution(end)
	if err != nil {
		return err
	}

	manager := manager.GetDataManager().ProcessInstanceDataManager
	err = manager.DeleteProcessInstance(end.Execution.GetProcessInstanceId())
	return err
}

func deleteDataForExecution(end *EndExecutionOperation) (err error) {
	taskManager := manager.GetDataManager().TaskDataManager
	tasks, errSelect := taskManager.FindByProcessInstanceId(end.Execution.GetTaskId())
	if errSelect == nil {
		for _, task := range tasks {
			taskManager.DeleteTask(task)
		}
	}
	processInstanceId := end.Execution.GetProcessInstanceId()
	identityLinkManager := manager.GetDataManager().IdentityLinkDataManager
	identityLinks, errSelect := identityLinkManager.SelectByProcessInstanceId(processInstanceId)
	if errSelect == nil {
		for _, identityLink := range identityLinks {
			err := identityLinkManager.Delete(identityLink.Id)
			if err != nil {
				return err
			}
		}
	}
	variableManager := manager.GetDataManager().VariableDataManager
	variables, err := variableManager.SelectByProcessInstanceId(processInstanceId)
	if err == nil {
		for _, variable := range variables {
			err = variableManager.Delete(variable.Id)
			if err != nil {
				return err
			}
		}
	}
	return err
}
