package cmd

import (
	"github.com/lios/go-activiti/engine/contanst"
	"github.com/lios/go-activiti/engine/event"
	"github.com/lios/go-activiti/engine/event/impl"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/handler"
	"github.com/lios/go-activiti/engine/impl/interceptor"
	. "github.com/lios/go-activiti/engine/impl/persistence/entity"
	"github.com/lios/go-activiti/engine/impl/utils"
	"github.com/lios/go-activiti/errs"
	"github.com/lios/go-activiti/logger"
)

type CompleteCmd struct {
	NeedsActiveTaskCmd
	Variables  map[string]interface{}
	LocalScope bool
}

func (completeCmd CompleteCmd) TaskExecute(command interceptor.CommandContext, entity TaskEntity) (interface{}, error) {
	err := completeCmd.executeTaskComplete(entity, command)
	return entity, err
}

func (completeCmd CompleteCmd) executeTaskComplete(entity TaskEntity, command interceptor.CommandContext) (err error) {

	// All properties set, now firing 'create' events
	if event.GetEventDispatcher().IsEnabled() {
		activitiEntityEvent, err := impl.CreateEntityEvent(event.TASK_COMPLETED, entity)
		if err != nil {
			return err
		}
		event.GetEventDispatcher().DispatchEvent(activitiEntityEvent)
	}
	err = GetTaskEntityManager().DeleteTask(entity)
	if err != nil {
		return err
	}
	task := entity.(*TaskEntityImpl)
	execution := GetExecutionEntityManager().FindById(task.ProcessInstanceId)
	if err != nil {
		return nil
	}
	processUtils := utils.ProcessDefinitionUtil{}
	process := processUtils.GetProcess(execution.GetProcessDefineId())
	currentTask := process.GetFlowElement(task.GetTaskDefineKey())
	task.SetCurrentFlowElement(currentTask)
	if completeCmd.LocalScope {
		//err = entity.SetVariable(entity, completeCmd.Variables)
	} else {
		err = entity.SetExecutionVariables(completeCmd.Variables)
	}
	if err != nil {
		return err
	}
	userTask, ok := currentTask.(*UserTask)
	if ok {
		taskListeners := userTask.ExtensionElements.TaskListener
		for _, listener := range taskListeners {
			if listener.EventType == contanst.TASK_TYPE_COMPLETED {
				err = handler.PerformTaskListener(&execution, userTask.Name, task.GetCurrentActivityId())
				if err != nil {
					return err
				}
			}
		}
	} else {
		logger.Error("not task")
		return errs.ProcessError{Code: "not task"}
	}
	command.Agenda.PlanTriggerExecutionOperation(task)
	return nil
}
