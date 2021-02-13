package deploy

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity"
	"github.com/lios/go-activiti/logger"
)

var (
	deploymentManager       DeploymentManager
	deploymentEntityManager DeploymentEntityManager
)

func init() {
	deploymentEntityManager = &DeploymentEntityManagerImpl{}
}

type DeploymentManager struct {
	Deployers []Deployer
}

func SetDeploymentManager(define DeploymentManager) {
	deploymentManager = define
}

func GetDeploymentManager() DeploymentManager {
	return deploymentManager
}
func (define DeploymentManager) Deploy(deployment DeploymentEntity, deploymentSettings map[string]interface{}) {
	for _, deployer := range define.Deployers {
		deployer.Deploy(deployment, deploymentSettings)
	}
}
func (define DeploymentManager) ResolveProcessDefinition(definitionEntity ProcessDefinitionEntity) ProcessDefinitionCacheEntry {
	deploymentId := definitionEntity.GetDeploymentId()
	deployment := deploymentEntityManager.FindById(deploymentId)
	deployment.SetNew(false)
	define.Deploy(deployment, nil)
	cacheEntry := ProcessDefinitionCacheEntry{}
	return cacheEntry
}

func (define DeploymentManager) FindDeployedProcessDefinitionById(processDefinitionId int64) ProcessDefinitionEntity {
	if processDefinitionId == 0 {
		logger.Error("processDefinitionId is nil")
		panic("processDefinitionId is nil")
	}
	manager := GetProcessDefinitionEntityManager()
	return manager.FindProcessDefinitionById(processDefinitionId)
}
