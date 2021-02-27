package behavior

import (
	. "github.com/lios/go-activiti/engine/contanst"
	"github.com/lios/go-activiti/engine/event"
	. "github.com/lios/go-activiti/engine/event/impl"
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
	. "github.com/lios/go-activiti/engine/impl/handler"
	"github.com/lios/go-activiti/engine/impl/interceptor"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
	. "github.com/lios/go-activiti/model"
	"time"
)

type UserTaskActivityBehavior struct {
	UserTask   model.UserTask
	ProcessKey string
}

//普通用户节点处理
func (user UserTaskActivityBehavior) Execute(execution delegate.DelegateExecution) (err error) {
	task := entity.TaskEntityImpl{}
	task.ProcessInstanceId = execution.GetProcessInstanceId()
	task.SetAssignee(user.UserTask.Assignee)

	task.SetStartTime(time.Now())
	task.SetTaskDefineKey(user.UserTask.Id)
	task.SetTaskDefineName(user.UserTask.Name)
	taskEntityManager := entity.GetTaskEntityManager()
	err = taskEntityManager.InsertTask(&task)
	if err != nil {
		return err
	}
	err = handleAssignments(user.UserTask, task.Id, task.ProcessInstanceId)

	// All properties set, now firing 'create' events
	if event.GetAtivitiEventDispatcher().IsEnabled() {
		activitiEntityEvent, err := CreateEntityEvent(event.TASK_CREATED, task)
		if err != nil {
			return err
		}
		event.GetAtivitiEventDispatcher().DispatchEvent(activitiEntityEvent)
	}
	extensionElements := user.UserTask.ExtensionElements
	if extensionElements.TaskListener != nil && len(extensionElements.TaskListener) > 0 {
		taskListeners := extensionElements.TaskListener
		for _, listener := range taskListeners {
			if listener.EventType == TASK_TYPE_CREATE {
				err = PerformTaskListener(execution, user.UserTask.Name, user.ProcessKey)
				if err != nil {
					return err
				}
			}
		}
	}
	return err
}

//保存候选用户
func handleAssignments(user model.UserTask, taskId, processInstanceId int64) (err error) {
	users := user.CandidateUsers
	if len(users) >= 0 {
		for _, user := range users {
			link := IdentityLink{}
			link.TaskId = taskId
			link.ProcessInstanceId = processInstanceId
			link.UserId = user
			identityLinkManager := entity.GetIdentityLinkManager()
			err = identityLinkManager.CreateIdentityLink(link)
			if err != nil {
				return err
			}
		}
	}
	return err
}

//普通用户节点处理
func (user UserTaskActivityBehavior) Trigger(execution delegate.DelegateExecution) {
	user.Leave(execution)
}

func (user UserTaskActivityBehavior) Leave(execution delegate.DelegateExecution) {
	element := execution.GetCurrentFlowElement()
	execution.SetCurrentFlowElement(element)
	interceptor.GetAgenda().Agenda.PlanTakeOutgoingSequenceFlowsOperation(execution, true)
}
