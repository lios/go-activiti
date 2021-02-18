package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	. "github.com/lios/go-activiti/model"
)

var (
	deploymentEntity      DeploymentEntity
	deploymentDataManager DeploymentDataManager
)

type DeploymentEntityManagerImpl struct {
	AbstractEntityManager
}

func init() {
	deploymentDataManager = DeploymentDataManager{AbstractDataManager: AbstractDataManager{TableModel{AbstractModel(Deployment{})}}}
}

func (deploymentEntity DeploymentEntityManagerImpl) GetDataManager() DataManagers {
	return deploymentDataManager
}

func (deploymentEntity DeploymentEntityManagerImpl) FindById(entityId int64) DeploymentEntity {
	manager := deploymentEntity.GetDataManager()
	dataManager := manager.(DeploymentDataManager)
	deployment := Deployment{}
	err := dataManager.FindById(entityId, &deployment)
	if err != nil {
		panic(err.Error())
	}
	deploymentEntityImpl := DeploymentEntityImpl{}
	deploymentEntityImpl.SetName(deployment.Name)
	deploymentEntityImpl.SetKey(deployment.Key)
	resourceDataManager := GetResourceDataManager().GetDataManager().(ResourceDataManager)
	resource, err := resourceDataManager.FindDeployedProcessDefinitionByDeploymentId(deployment.Id)
	if err != nil {
		panic(err)
	}
	resourceEntityImpl := ResourceEntityImpl{}
	resourceEntityImpl.SetName(resource.Name)
	resourceEntityImpl.SetDeploymentId(resource.DeploymentId)
	resourceEntityImpl.SetBytes([]byte(resource.Bytes))
	deploymentEntityImpl.SetResources(&resourceEntityImpl)
	return &deploymentEntityImpl
}
