package entity

type ResourceEntityManager interface {
	FindResourcesByDeploymentId(deploymentId int64) ResourceEntity
}
