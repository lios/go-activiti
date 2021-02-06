package entity

type ExecutionEntityManager interface {
	FindById(entityId int64) ExecutionEntity
}
