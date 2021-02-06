package entity

import (
	. "github.com/lios/go-activiti/engine/variable"
)

type TaskEntityImpl struct {
	AbstractEntity
	ExecutionEntityImpl
	VariableScopeImpl
	TaskId      int64
	TaskName    string
	Variables   map[string]interface{}
	executionId int64
	execution   *ExecutionEntityImpl
}

func (taskEntiy TaskEntityImpl) GetTaskName() string {
	return taskEntiy.TaskName
}

func (taskEntiy TaskEntityImpl) GetById(id int64) Entity {
	taskEntityImpl := TaskEntityImpl{}
	task, err := taskDataManager.GetById(id)
	if err != nil {
		return taskEntityImpl
	}
	taskEntityImpl.TaskName = task.TaskDefineName
	return taskEntityImpl
}

func (taskEntiy TaskEntityImpl) GetTaskId() int64 {
	return taskEntiy.TaskId
}

func (taskEntiy TaskEntityImpl) SetTaskId(taskId int64) {
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
	//variableManager := task.GetVariableEntityManager()
	//return variableManager.SelectTaskId(variableName, taskEntiy.TaskId)
	return Variable{}, nil
}
func (taskEntiy TaskEntityImpl) getExecution() ExecutionEntity {
	if taskEntiy.execution == nil && taskEntiy.executionId != 0 {
		executionEntityManager := GetExecutionEntityManager()
		entityImpl := executionEntityManager.FindById(taskEntiy.executionId)
		taskEntiy.execution = &entityImpl
	}
	return nil
}
func (taskEntiy TaskEntityImpl) SetExecutionVariables(variables map[string]interface{}) error {
	if taskEntiy.getExecution() != nil {
		taskEntiy.execution.SetVariableLocal(variables)
	}
	return nil
}
func (taskEntiy TaskEntityImpl) SetScope(variable *Variable) {
	variable.TaskId = taskEntiy.TaskId
}
