package cmd

import (
	"github.com/lios/go-activiti/engine/common"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/handler"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
	"github.com/lios/go-activiti/engine/interceptor"
	"github.com/lios/go-activiti/event"
	"github.com/lios/go-activiti/event/impl"
	"github.com/lios/go-activiti/model"
)

type CompleteCmd struct {
	TaskId     int
	Variables  map[string]interface{}
	LocalScope bool
}

func (taskCmd CompleteCmd) Execute(command interceptor.CommandContext) (interface{}, error) {
	manager := command.ProcessEngineConfiguration.TaskDataManager
	task, err := manager.FindById(taskCmd.TaskId)
	if err != nil {
		return task, err
	}
	taskCmd.executeTaskComplete(task, command)
	return task, err
}

func (taskCmd CompleteCmd) executeTaskComplete(task model.Task, command interceptor.CommandContext) (err error) {

	// All properties set, now firing 'create' events
	if event.GetEventDispatcher().IsEnabled() {
		activitiEntityEvent, err := impl.CreateEntityEvent(event.TASK_COMPLETED, task)
		if err != nil {
			return err
		}
		event.GetEventDispatcher().DispatchEvent(activitiEntityEvent)
	}
	err = deleteTask(task, command)
	if err != nil {
		return err
	}
	//defineManager := command.ProcessEngineConfiguration.DefineDataManager
	//bytearry, err := defineManager.FindProcessByTask(task.ProcessInstanceId)
	//if err != nil {
	//	return err
	//}
	taskExecution := command.ProcessEngineConfiguration.ExecutionEntityManager.FindById(task.TaskDefineKey)
	currentTask := taskExecution.GetCurrentFlowElement()
	//currentTask := behavior.FindCurrentTask(bytearry, task.TaskDefineKey)
	//taskExecution := entity.TaskEntityImpl{}
	execution := entity.ExecutionEntityImpl{}
	execution.SetCurrentFlowElement(currentTask)
	execution.SetCurrentActivityId(task.TaskDefineKey)
	processInstanceManager := interceptor.GetProcessEngineConfiguration().ProcessInstanceDataManager
	instance := processInstanceManager.GetProcessInstance(task.ProcessInstanceId)
	execution.SetProcessDefineId(instance.ProcessDefineId)
	execution.SetProcessInstanceId(task.ProcessInstanceId)
	taskExecution.SetTaskId(task.Id)
	//taskExecution.ExecutionEntityImpl = execution
	if taskCmd.LocalScope {
		err = entity.SetVariable(taskExecution, taskCmd.Variables)
	} else {
		err = entity.SetVariable(&execution, taskCmd.Variables)
	}
	if err != nil {
		return err
	}

	userTask, ok := currentTask.(UserTask)
	if ok {
		taskListeners := userTask.ExtensionElements.TaskListener
		for _, listener := range taskListeners {
			if listener.EventType == common.TASK_TYPE_COMPLETED {
				err = handler.PerformTaskListener(&execution, userTask.Name, instance.Key)
				if err != nil {
					return err
				}
			}
		}
	}
	command.Agenda.PlanTriggerExecutionOperation(taskExecution)
	return nil
}

func deleteTask(task model.Task, command interceptor.CommandContext) (err error) {
	manager := command.ProcessEngineConfiguration.TaskDataManager
	return manager.DeleteTask(task)
}
