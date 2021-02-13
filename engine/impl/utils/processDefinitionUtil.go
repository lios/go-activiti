package utils

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/persistence/deploy"
)

type ProcessDefinitionUtil struct {
}

func (ProcessDefinitionUtil) GetProcess(processDefinitionId int64) model.Process {
	deploymentManager := deploy.GetDeploymentManager()
	definitionEntity := deploymentManager.FindDeployedProcessDefinitionById(processDefinitionId)
	return deploymentManager.ResolveProcessDefinition(definitionEntity).Process
}
