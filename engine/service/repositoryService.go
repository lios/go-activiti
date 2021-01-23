package engine

import (
	"github.com/lios/go-activiti/engine/behavior"
	"github.com/lios/go-activiti/engine/cmd"
)

type RepositoryService struct {
}

//流程部署
func (this RepositoryService) Deploy(name string, key string, bytes []byte) {
	behavior.GetServiceImpl().CommandExecutor.Exe(cmd.DeploymentCmd{Name: name, Key: key, Bytes: bytes})
}
