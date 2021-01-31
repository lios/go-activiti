package parse

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type ParsedDeploymentBuilder struct {
	deployment         entity.DeploymentEntity
	bpmnParser         *BpmnParse
	deploymentSettings map[string]interface{}
}

func NewParsedDeploymentBuilder(deployment entity.DeploymentEntity, bpmnParser *BpmnParse, deploymentSettings map[string]interface{}) ParsedDeploymentBuilder {
	return ParsedDeploymentBuilder{deployment, bpmnParser, deploymentSettings}
}

func (parsedDeploymentBuilder ParsedDeploymentBuilder) Build() ParsedDeployment {
	resources := parsedDeploymentBuilder.deployment.GetResources()
	for _, resource := range resources {
		parsedDeploymentBuilder.createBpmnParseFromResource(resource)
	}
	return ParsedDeployment{parsedDeploymentBuilder.bpmnParser}
}

func (parsedDeploymentBuilder ParsedDeploymentBuilder) createBpmnParseFromResource(resource entity.ResourceEntity) {
	parsedDeploymentBuilder.bpmnParser.Execute()
}
