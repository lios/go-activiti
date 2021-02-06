package entity

type ProcessDefinitionEntity interface {
	SetKey(key string)

	GetKey() string

	SetName(name string)

	GetName() string

	SetDescription(description string)

	GetDescription() string

	SetDeploymentId(deploymentId int64)

	GetDeploymentId() int64

	SetVersion(version int)

	GetVersion() int

	SetResourceName(resourceName string)

	GetResourceName() string
}
