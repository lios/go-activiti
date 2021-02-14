package entity

import "time"

type DeploymentEntity interface {
	AddResource(resource ResourceEntity)

	GetResources() ResourceEntity

	GetName() string

	SetName(name string)

	SetKey(key string)

	GetKey() string

	SetTenantId(tenantId string)

	SetResources(ResourceEntity)

	SetDeploymentTime(deploymentTime time.Time)

	New() bool

	SetNew(boolean bool)
}
