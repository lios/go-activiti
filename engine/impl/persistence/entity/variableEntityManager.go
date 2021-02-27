package entity

type VariableEntityManager interface {
	CreteVariable() error
	DeleteVariableInstanceByTask(taskId int64)
}
