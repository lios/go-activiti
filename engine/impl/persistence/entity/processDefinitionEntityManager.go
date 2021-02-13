package entity

type ProcessDefinitionEntityManager interface {
	FindProcessDefinitionById(processDefinitionId int64) ProcessDefinitionEntity
	FindLatestProcessDefinitionByKey(processDefinitionKey string) (ProcessDefinitionEntity, error)
}
