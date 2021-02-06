package entity

type AbstractEntityManager struct {
	EntityManager

	//GetTaskEntityManager() TaskDataManager
	//
	//GetDefineEntityManager() DefineDataManager
	//
	//GetDeploymentDataManager() DeploymentDataManager
	//
	//GetVariableEntityManager() VariableDataManager
	//
	//GetIdentityLinkEntityManager() IdentityLinkDataManager
	//
	//GetHistoricActinstEntityManager() HistoricActinstDataManager
	//
	//GetHistoricTaskEntityManager() HistoricTaskDataManager
	//
	//GetHistoricProcessEntityManager() HistoricProcessDataManager
	//
	//GetResourceDataManager() ResourceDataManager
}

func (entityManager AbstractEntityManager) Insert(data interface{}) error {
	return entityManager.GetDataManager().Insert(data)
}

func (entityManager AbstractEntityManager) GetById(id int64) Entity {
	return AbstractEntity{}
}

func (entityManager AbstractEntityManager) Delete(entity Entity) {
	entityManager.GetDataManager().Delete(entity.GetId())
}
