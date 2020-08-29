package engine

type Runtime interface {
	//发起流程
	StartProcessInstanceByKey(processDefinitionKey string, variables map[string]interface{}, businessKey string, tenantId string)
}
