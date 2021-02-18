package behavior

import (
	. "github.com/lios/go-activiti/engine/contanst"
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	. "github.com/lios/go-activiti/engine/impl/handler"
	"github.com/lios/go-activiti/engine/impl/interceptor"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	"reflect"
	"time"
)

type UserAutoTaskActivityBehavior struct {
	UserTask   model.UserTask
	ProcessKey string
}

//自动通过用户节点处理
func (user UserAutoTaskActivityBehavior) Execute(execution entity.ExecutionEntity) (err error) {
	task := entity.TaskEntityImpl{}
	task.ProcessInstanceId = execution.GetProcessInstanceId()
	task.SetAssignee(user.UserTask.Assignee)
	task.SetStartTime(time.Now())
	task.SetTaskDefineKey(user.UserTask.Id)
	task.SetTaskDefineName(user.UserTask.Name)
	dataManager := entity.GetTaskEntityManager()
	taskDataManager := dataManager.GetDataManager().(TaskDataManager)
	err = taskDataManager.Insert(task)

	activitiConstructor, err := GetConstructorByName(user.ProcessKey)
	if err != nil {
		dataManager.DeleteTask(&task)
		interceptor.GetAgenda().Agenda.PlanTriggerExecutionOperation(execution)
		return nil
	}
	constructor := activitiConstructor(execution)
	reflectConstructor := reflect.ValueOf(constructor)
	taskParams := []reflect.Value{reflectConstructor}

	method, b := reflectConstructor.Type().MethodByName(user.UserTask.Name)
	if !b {
		dataManager.DeleteTask(&task)
		interceptor.GetAgenda().Agenda.PlanTriggerExecutionOperation(execution)
		return err
	}

	callResponse := method.Func.Call(taskParams)

	code := callResponse[0].Interface()
	errRes := callResponse[1].Interface()
	code = code.(string)
	if code != ACTIVITI_HANDLER_CODE {
		err := errRes.(error)
		return err
	}
	dataManager.DeleteTask(&task)
	interceptor.GetAgenda().Agenda.PlanTriggerExecutionOperation(execution)
	return err
}

//普通用户节点处理
func (user UserAutoTaskActivityBehavior) Trigger(execution entity.ExecutionEntity) {
	user.Leave(execution)
}

func (user UserAutoTaskActivityBehavior) Leave(execution entity.ExecutionEntity) {
	element := execution.GetCurrentFlowElement()
	execution.SetCurrentFlowElement(element)
	interceptor.GetAgenda().Agenda.PlanTakeOutgoingSequenceFlowsOperation(execution, true)
}
