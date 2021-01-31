package engine

type RuntimeService interface {
	StartProcessInstanceByKey(processDefinitionKey string, variables map[string]interface{}, businessKey string, tenantId string)
}
