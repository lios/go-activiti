package behavior

import (
	"github.com/lios/go-activiti/engine"
	. "github.com/lios/go-activiti/engine/common"
	. "github.com/lios/go-activiti/engine/handler"
	. "github.com/lios/go-activiti/engine/persistence"
	. "github.com/lios/go-activiti/model"
	"reflect"
	"time"
)

type UserAutoTaskActivityBehavior struct {
	UserTask   engine.UserTask
	ProcessKey string
}

//自动通过用户节点处理
func (user UserAutoTaskActivityBehavior) Execute(execution engine.ExecutionEntity) (err error) {
	task := Task{}
	task.ProcessInstanceId = execution.GetProcessInstanceId()
	task.Assignee = user.UserTask.Assignee
	task.StartTime = time.Now()
	task.TaskDefineKey = user.UserTask.Id
	task.TaskDefineName = user.UserTask.Name
	manager := TaskManager{Task: &task}
	err = manager.Insert(execution)

	activitiConstructor, err := GetConstructorByName(user.ProcessKey)
	if err != nil {
		return err
	}
	constructor := activitiConstructor(execution)
	reflectConstructor := reflect.ValueOf(constructor)
	taskParams := []reflect.Value{reflectConstructor}

	method, b := reflectConstructor.Type().MethodByName(user.UserTask.Name)
	if !b {
		manager.DeleteTask(task)
		GetAgenda().PlanTriggerExecutionOperation(execution)
	}

	callResponse := method.Func.Call(taskParams)

	code := callResponse[0].Interface()
	errRes := callResponse[1].Interface()
	code = code.(string)
	if code != ACTIVITI_HANDLER_CODE {
		err := errRes.(error)
		return err
	}
	manager.DeleteTask(task)
	GetAgenda().PlanTriggerExecutionOperation(execution)
	return err
}

//普通用户节点处理
func (user UserAutoTaskActivityBehavior) Trigger(execution engine.ExecutionEntity) {
	user.Leave(execution)
}

func (user UserAutoTaskActivityBehavior) Leave(execution engine.ExecutionEntity) {
	element := execution.GetCurrentFlowElement()
	execution.SetCurrentFlowElement(element)
	GetAgenda().PlanTakeOutgoingSequenceFlowsOperation(execution, true)
}
