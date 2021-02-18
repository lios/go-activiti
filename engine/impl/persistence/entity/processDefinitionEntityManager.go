package entity

type ProcessDefinitionEntityManager interface {
	EntityManager
	FindProcessDefinitionById(processDefinitionId int64) ProcessDefinitionEntity
	FindLatestProcessDefinitionByKey(processDefinitionKey string) (ProcessDefinitionEntity, error)
}
