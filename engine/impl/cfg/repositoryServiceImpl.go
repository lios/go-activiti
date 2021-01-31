package cfg

import (
	"github.com/lios/go-activiti/engine/impl/cmd"
)

type RepositoryServiceImpl struct {
	ServiceImpl
}

//流程部署
func (this RepositoryServiceImpl) Deploy(name string, key string, bytes []byte) {
	this.GetCommandExecutor().Exe(cmd.DeploymentCmd{Name: name, Key: key, Bytes: bytes})
}
