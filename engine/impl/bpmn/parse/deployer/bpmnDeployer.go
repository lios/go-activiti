package deployer

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/bpmn/parse"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type BpmnDeployer struct {
	ParsedDeployment               parse.ParsedDeployment
	ParsedDeploymentBuilderFactory parse.ParsedDeploymentBuilderFactory
}

func (bpmnDeployer *BpmnDeployer) Deploy(deploymentEntity entity.DeploymentEntity, deploymentSettings map[string]interface{}) {
	parsedDeployment := bpmnDeployer.ParsedDeploymentBuilderFactory.GetBuilderForDeploymentAndSettings(deploymentEntity, deploymentSettings).Build()
	bpmnDeployer.ParsedDeployment = parsedDeployment
}
func (bpmnDeployer BpmnDeployer) GetProcess(key string) model.Process {
	return *bpmnDeployer.ParsedDeployment.BpmnParse.BpmnModel.GetProcessById(key)
}
func (bpmnDeployer BpmnDeployer) SetParsedDeploymentBuilderFactory(parsedDeploymentBuilderFactory parse.ParsedDeploymentBuilderFactory) {
	bpmnDeployer.ParsedDeploymentBuilderFactory = parsedDeploymentBuilderFactory
}
