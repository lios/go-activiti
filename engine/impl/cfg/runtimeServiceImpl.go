package cfg

import (
	"github.com/lios/go-activiti/engine/impl/cmd"
)

type RuntimeServiceImpl struct {
	ServiceImpl
}

//发起流程
func (runtime RuntimeServiceImpl) StartProcessInstanceByKey(processDefinitionKey string, variables map[string]interface{},
	businessKey string, tenantId string) {
	runtime.GetCommandExecutor().Exe(cmd.StartProcessInstanceByKeyCmd{ProcessDefinitionKey: processDefinitionKey,
		Variables: variables, TenantId: tenantId, BusinessKey: businessKey})
}
