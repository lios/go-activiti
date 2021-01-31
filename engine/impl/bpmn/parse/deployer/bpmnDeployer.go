package deployer

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/parse"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type BpmnDeployer struct {
	ParsedDeploymentBuilderFactory parse.ParsedDeploymentBuilderFactory
}

func (bpmnDeployer BpmnDeployer) Deploy(deploymentEntity entity.DeploymentEntity, deploymentSettings map[string]interface{}) {
	bpmnDeployer.ParsedDeploymentBuilderFactory.GetBuilderForDeploymentAndSettings(deploymentEntity, deploymentSettings).Build()
}

func (bpmnDeployer BpmnDeployer) SetParsedDeploymentBuilderFactory(parsedDeploymentBuilderFactory parse.ParsedDeploymentBuilderFactory) {
	bpmnDeployer.ParsedDeploymentBuilderFactory = parsedDeploymentBuilderFactory
}
