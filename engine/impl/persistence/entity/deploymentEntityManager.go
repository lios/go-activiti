package entity

type DeploymentEntityManager interface {
	FindById(entityId int64) DeploymentEntity
}
