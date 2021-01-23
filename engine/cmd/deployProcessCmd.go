package cmd

import (
	"github.com/lios/go-activiti/engine/behavior"
	"github.com/lios/go-activiti/engine/persistence"
)

type DeploymentCmd struct {
	Name     string
	Key      string
	TenantId string
	Bytes    []byte
}

func (deploy DeploymentCmd) Execute(interceptor behavior.CommandContext) (interface{}, error) {
	deploymentManager := persistence.DeploymentManager{}
	err := deploymentManager.Deployment(deploy.Name, deploy.Key, deploy.Bytes)
	return nil, err
}
