package deploy

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type Deployer interface {
	Deploy(deploymentEntity entity.DeploymentEntity, deploymentSettings map[string]interface{})
	GetProcess(key string) model.Process
}
