package cmd

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	"github.com/lios/go-activiti/engine/interceptor"
)

type DeploymentCmd struct {
	Name     string
	Key      string
	TenantId string
	Bytes    []byte
}

func (deploy DeploymentCmd) Execute(interceptor interceptor.CommandContext) (interface{}, error) {
	deploymentManager := data.DeploymentDataManager{}
	err := deploymentManager.Deployments(deploy.Name, deploy.Key, deploy.Bytes)
	return nil, err
}
