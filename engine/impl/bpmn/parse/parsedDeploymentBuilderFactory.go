package parse

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type ParsedDeploymentBuilderFactory struct {
	BpmnParser *BpmnParser
}

func (parsedDeploymentBuilderFactory ParsedDeploymentBuilderFactory) GetBuilderForDeploymentAndSettings(deployment entity.DeploymentEntity, deploymentSettings map[string]interface{}) ParsedDeploymentBuilder {
	return ParsedDeploymentBuilder{deployment, parsedDeploymentBuilderFactory.BpmnParser, deploymentSettings}
}
