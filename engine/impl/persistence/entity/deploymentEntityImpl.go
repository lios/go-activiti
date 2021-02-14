package entity

import (
	"time"
)

type DeploymentEntityImpl struct {
	Name           string
	Key            string
	TenantId       string
	DeploymentTime time.Time
	IsNew          bool
	ResourceEntity ResourceEntity
}

func (deploymentEntity *DeploymentEntityImpl) AddResource(resource ResourceEntity) {
	deploymentEntity.ResourceEntity = resource
}

func (deploymentEntity *DeploymentEntityImpl) GetResources() ResourceEntity {
	return deploymentEntity.ResourceEntity
}

func (deploymentEntity *DeploymentEntityImpl) GetName() string {
	return deploymentEntity.Name
}

func (deploymentEntity *DeploymentEntityImpl) SetName(name string) {

}

func (deploymentEntity *DeploymentEntityImpl) SetKey(key string) {
	deploymentEntity.Key = key
}
func (deploymentEntity *DeploymentEntityImpl) GetKey() string {
	return deploymentEntity.Key
}
func (deploymentEntity *DeploymentEntityImpl) SetTenantId(tenantId string) {

}

func (deploymentEntity *DeploymentEntityImpl) SetResources(resourceEntity ResourceEntity) {
	deploymentEntity.ResourceEntity = resourceEntity
}

func (deploymentEntity *DeploymentEntityImpl) SetDeploymentTime(deploymentTime time.Time) {

}

func (deploymentEntity *DeploymentEntityImpl) New() bool {
	return deploymentEntity.IsNew
}

func (deploymentEntity *DeploymentEntityImpl) SetNew(boolean bool) {
	deploymentEntity.IsNew = boolean
}
