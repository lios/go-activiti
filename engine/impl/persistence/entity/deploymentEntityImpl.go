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
	ResourceEntity map[string]ResourceEntity
}

func (deploymentEntity DeploymentEntityImpl) AddResource(resource ResourceEntity) {
	deploymentEntity.ResourceEntity[GetName()] = resource
}

func (deploymentEntity DeploymentEntityImpl) GetResources() map[string]ResourceEntity {
	return deploymentEntity.ResourceEntity
}

func (deploymentEntity DeploymentEntityImpl) GetName() string {
	return deploymentEntity.Name
}

func (deploymentEntity DeploymentEntityImpl) SetName(name string) {

}

func (deploymentEntity DeploymentEntityImpl) SetKey(key string) {

}

func (deploymentEntity DeploymentEntityImpl) SetTenantId(tenantId string) {

}

func (deploymentEntity DeploymentEntityImpl) SetResources(map[string]ResourceEntity) {

}

func (deploymentEntity DeploymentEntityImpl) SetDeploymentTime(deploymentTime time.Time) {

}

func (deploymentEntity DeploymentEntityImpl) New() bool {
	return deploymentEntity.IsNew
}

func (deploymentEntity DeploymentEntityImpl) SetNew(boolean bool) {

}
