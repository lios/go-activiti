package entity

type VariableEntityManager interface {
	DeleteVariableInstanceByTask(taskId int64)
}
