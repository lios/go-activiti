package entity

type ExecutionEntityManager interface {
	FindById(entityId string) ExecutionEntity
}
