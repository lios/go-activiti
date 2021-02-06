package entity

type ResourceEntity interface {
	GetName() string

	SetName(name string)

	GetBytes() []byte

	SetBytes(bytes []byte)

	GetDeploymentId() int64

	SetDeploymentId(deploymentId int64)
}
