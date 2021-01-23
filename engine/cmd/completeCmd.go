package cmd

import (
	"github.com/lios/go-activiti/engine/behavior"
	. "github.com/lios/go-activiti/engine/entityImpl"
	. "github.com/lios/go-activiti/engine/manager"
	"github.com/lios/go-activiti/event"
	"github.com/lios/go-activiti/event/impl"
	"github.com/lios/go-activiti/model"
)

type CompleteCmd struct {
	TaskId     int
	Variables  map[string]interface{}
	LocalScope bool
}

func (taskCmd CompleteCmd) Execute(interceptor behavior.CommandContext) (interface{}, error) {
	manager := GetTaskManager()
	task, err := manager.FindById(taskCmd.TaskId)
	if err != nil {
		return task, err
	}
	taskCmd.executeTaskComplete(task, interceptor)
	return task, err
}

func (taskCmd CompleteCmd) executeTaskComplete(task model.Task, interceptor behavior.CommandContext) (err error) {

	// All properties set, now firing 'create' events
	if event.GetEventDispatcher().IsEnabled() {
		activitiEntityEvent, err := impl.CreateEntityEvent(event.TASK_COMPLETED, task)
		if err != nil {
			return err
		}
		event.GetEventDispatcher().DispatchEvent(activitiEntityEvent)
	}
	err = deleteTask(task)
	if err != nil {
		return err
	}
	defineManager := GetDefineManager()
	bytearry, err := defineManager.FindProcessByTask(task.ProcessInstanceId)
	if err != nil {
		return err
	}
	currentTask := behavior.FindCurrentTask(bytearry, task.TaskDefineKey)
	taskExecution := TaskEntityImpl{}
	execution := ExecutionEntityImpl{}
	execution.SetCurrentFlowElement(currentTask)
	execution.SetCurrentActivityId(task.TaskDefineKey)
	processInstanceManager := behavior.GetProcessInstanceManager()
	execution.SetProcessDefineId(processInstanceManager.GetProcessInstance(task.ProcessInstanceId).ProcessDefineId)
	execution.SetProcessInstanceId(task.ProcessInstanceId)
	taskExecution.SetTaskId(task.Id)
	taskExecution.ExecutionEntityImpl = execution
	if taskCmd.LocalScope {
		err = SetVariable(&taskExecution, taskCmd.Variables)
	} else {
		err = SetVariable(&execution, taskCmd.Variables)
	}
	if err != nil {
		return err
	}
	interceptor.Agenda.PlanTriggerExecutionOperation(&taskExecution)
	return nil
}

func deleteTask(task model.Task) (err error) {
	manager := GetTaskManager()
	return manager.DeleteTask(task)
}
