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
	ProcessDefinitionCacheEntry ProcessDefinitionCacheEntry
	Deployers                   []Deployer
}

func SetDeploymentManager(define DeploymentManager) {
	deploymentManager = define
}

func GetDeploymentManager() DeploymentManager {
	return deploymentManager
}
func (define *DeploymentManager) Deploys(deployment DeploymentEntity, deploymentSettings map[string]interface{}) {
	var cacheEntry ProcessDefinitionCacheEntry
	for _, deployer := range define.Deployers {
		deployer.Deploy(deployment, deploymentSettings)
		process := deployer.GetProcess(deployment.GetKey())
		cacheEntry = NewProcessDefinitionCacheEntry(process)
	}
	define.ProcessDefinitionCacheEntry = cacheEntry
}
func (define *DeploymentManager) ResolveProcessDefinition(definitionEntity ProcessDefinitionEntity) ProcessDefinitionCacheEntry {
	deploymentId := definitionEntity.GetDeploymentId()
	deployment := deploymentEntityManager.FindById(deploymentId)
	deployment.SetNew(false)
	define.Deploys(deployment, nil)
	return define.ProcessDefinitionCacheEntry
}

func (define DeploymentManager) FindDeployedProcessDefinitionById(processDefinitionId int64) ProcessDefinitionEntity {
	if processDefinitionId == 0 {
		logger.Error("processDefinitionId is nil")
		panic("processDefinitionId is nil")
	}
	manager := GetProcessDefinitionEntityManager()
	return manager.FindProcessDefinitionById(processDefinitionId)
}
