package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	. "github.com/lios/go-activiti/engine/variable"
	"time"
)

type TaskEntityImpl struct {
	AbstractEntity
	ExecutionEntityImpl
	VariableScopeImpl
	TaskId         int64
	TaskName       string
	assignee       string
	startTime      time.Time
	taskDefineKey  string
	taskDefineName string
	Variables      map[string]interface{}
	executionId    int64
	execution      *ExecutionEntityImpl
}

func (taskEntiy TaskEntityImpl) GetTaskName() string {
	return taskEntiy.TaskName
}
func (taskEntiy *TaskEntityImpl) SetAssignee(assignee string) {
	taskEntiy.assignee = assignee
}

func (taskEntiy *TaskEntityImpl) SetStartTime(startTime time.Time) {
	taskEntiy.startTime = startTime
}

func (taskEntiy *TaskEntityImpl) SetTaskDefineKey(taskDefineKey string) {
	taskEntiy.taskDefineKey = taskDefineKey
}

func (taskEntiy *TaskEntityImpl) GetTaskDefineKey() string {
	return taskEntiy.taskDefineKey
}

func (taskEntiy *TaskEntityImpl) SetTaskDefineName(taskDefineName string) {
	taskEntiy.taskDefineName = taskDefineName
}

func (taskEntiy TaskEntityImpl) GetTaskId() int64 {
	return taskEntiy.TaskId
}

func (taskEntiy *TaskEntityImpl) SetTaskId(taskId int64) {
	taskEntiy.TaskId = taskId
}

func (taskEntiy TaskEntityImpl) GetVariable() map[string]interface{} {
	//variableManager := task.GetVariableEntityManager()
	//variables, err := variableManager.SelectByTaskId(taskEntiy.TaskId)
	//if err != nil {
	//	return task.HandleVariable(variables)
	//}
	return nil
}

func (taskEntiy TaskEntityImpl) GetSpecificVariable(variableName string) (Variable, error) {
	manager := GetVariableEntityManager().GetDataManager()
	variableDataManager := manager.(VariableDataManager)
	return variableDataManager.SelectTaskId(variableName, taskEntiy.TaskId)
}
func (taskEntiy *TaskEntityImpl) getExecution() ExecutionEntity {
	if taskEntiy.execution == nil {
		executionEntityManager := GetExecutionEntityManager()
		entityImpl := executionEntityManager.FindById(taskEntiy.ProcessInstanceId)
		taskEntiy.execution = &entityImpl
		return &entityImpl
	}
	return nil
}
func (taskEntiy *TaskEntityImpl) SetExecutionVariables(variables map[string]interface{}) error {
	return taskEntiy.SetVariable(taskEntiy, variables)
}
func (taskEntiy *TaskEntityImpl) SetScope(variable *Variable) {
	variable.TaskId = taskEntiy.TaskId
}
