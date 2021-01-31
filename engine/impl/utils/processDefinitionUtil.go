package utils

import (
	"github.com/lios/go-activiti/engine"
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type ProcessDefinitionUtil struct {
}

func (ProcessDefinitionUtil) GetProcess(processDefinitionId string) model.Process {
	deploymentManager := engine.GetProcessEngineConfiguration().DeploymentManager
	deployment := deploymentManager.FindDeployedProcessDefinitionById(processDefinitionId)
	return deploymentManager.ResolveProcessDefinition(deployment).Process
}
