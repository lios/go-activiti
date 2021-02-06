package parse

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type ParsedDeploymentBuilder struct {
	deployment         entity.DeploymentEntity
	bpmnParser         *BpmnParser
	deploymentSettings map[string]interface{}
}

func NewParsedDeploymentBuilder(deployment entity.DeploymentEntity, bpmnParser *BpmnParser, deploymentSettings map[string]interface{}) ParsedDeploymentBuilder {
	return ParsedDeploymentBuilder{deployment, bpmnParser, deploymentSettings}
}

func (parsedDeploymentBuilder ParsedDeploymentBuilder) Build() ParsedDeployment {
	resources := parsedDeploymentBuilder.deployment.GetResources()
	bpmnParse := parsedDeploymentBuilder.createBpmnParseFromResource(resources)
	return ParsedDeployment{BpmnParse: bpmnParse}
}

func (parsedDeploymentBuilder ParsedDeploymentBuilder) createBpmnParseFromResource(resource entity.ResourceEntity) BpmnParse {
	name := resource.GetName()
	bytes := resource.GetBytes()
	bpmnParse := parsedDeploymentBuilder.bpmnParser.CreateParse().SourceInputStream(bytes).Deployment(parsedDeploymentBuilder.deployment).SourceName(name)
	bpmnParse.Execute()
	return bpmnParse
}
