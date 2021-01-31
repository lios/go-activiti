package entity

import (
	. "github.com/lios/go-activiti/engine/variable"
)

type TaskEntityImpl struct {
	AbstractEntityManager
	ExecutionEntityImpl
	TaskId    int64
	Variables map[string]interface{}
}

func (task *TaskEntityImpl) GetTaskId() int64 {
	return task.TaskId
}

func (task *TaskEntityImpl) SetTaskId(taskId int64) {
	task.TaskId = taskId
}

func (task *TaskEntityImpl) GetVariable() map[string]interface{} {
	variableManager := task.GetVariableEntityManager()
	variables, err := variableManager.SelectByTaskId(task.TaskId)
	if err != nil {
		return task.HandleVariable(variables)
	}
	return nil
}

func (task *TaskEntityImpl) GetSpecificVariable(variableName string) (Variable, error) {
	variableManager := task.GetVariableEntityManager()
	return variableManager.SelectTaskId(variableName, task.TaskId)
}

func (task *TaskEntityImpl) SetScope(variable *Variable) {
	variable.TaskId = task.TaskId
}
