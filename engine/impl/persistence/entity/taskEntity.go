package entity

type TaskEntity interface {
	Entity
	GetTaskName() string

	SetExecutionVariables(parameters map[string]interface{}) error
}
