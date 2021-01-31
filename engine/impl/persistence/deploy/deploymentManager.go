package deploy

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type DeploymentManager struct {
	deployers []Deployer
}

func (define DeploymentManager) Deploy(deployment entity.DeploymentEntity, deploymentSettings map[string]interface{}) {
	for _, deployer := range define.deployers {
		Deploy(deployment, deploymentSettings)
	}
}
func (define DeploymentManager) ResolveProcessDefinition(deployment entity.ProcessInstanceEntity) ProcessDefinitionCacheEntry {
	deploymentEntity := entity.DeploymentEntityImpl{}
	define.Deploy(deploymentEntity, nil)
	cacheEntry := ProcessDefinitionCacheEntry{}
	return cacheEntry
}

func (define DeploymentManager) FindDeployedProcessDefinitionById(processDefinitionId string) entity.ProcessDefinitionEntityImpl {
	return entity.ProcessDefinitionEntityImpl{}
}
