package entity

import "time"

type DeploymentEntity interface {
	AddResource(resource ResourceEntity)

	GetResources() map[string]ResourceEntity

	GetName() string

	SetName(name string)

	SetKey(key string)

	SetTenantId(tenantId string)

	SetResources(map[string]ResourceEntity)

	SetDeploymentTime(deploymentTime time.Time)

	New() bool

	SetNew(boolean bool)
}
