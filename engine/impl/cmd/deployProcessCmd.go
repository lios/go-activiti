package cmd

import (
	"github.com/lios/go-activiti/engine/impl/interceptor"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
	"github.com/lios/go-activiti/engine/impl/persistence/entity/data"
)

type DeploymentCmd struct {
	Name     string
	Key      string
	TenantId string
	Bytes    []byte
}

func (deploy DeploymentCmd) Execute(interceptor interceptor.CommandContext) (interface{}, error) {
	deploymentManager := entity.GetDeploymentEntityManager().GetDataManager()
	deploymentDataManager := deploymentManager.(data.DeploymentDataManager)
	err := deploymentDataManager.Deployments(deploy.Name, deploy.Key, deploy.Bytes)
	return nil, err
}
