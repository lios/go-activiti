package engine

import (
	"github.com/lios/go-activiti/engine/behavior"
	"github.com/lios/go-activiti/engine/cmd"
)

type RuntimeService struct {
	behavior.ServiceImpl
}

//发起流程
func (runtime RuntimeService) StartProcessInstanceByKey(processDefinitionKey string, variables map[string]interface{},
	businessKey string, tenantId string) {
	behavior.GetServiceImpl().CommandExecutor.Exe(cmd.StartProcessInstanceByKeyCmd{ProcessDefinitionKey: processDefinitionKey,
		Variables: variables, TenantId: tenantId, BusinessKey: businessKey})
}
