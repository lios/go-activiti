package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	. "github.com/lios/go-activiti/model"
)

var executionDataManager ExecutionDataManager

type ExecutionEntityManagerImpl struct {
	AbstractEntityManager
}

func init() {
	executionDataManager = ExecutionDataManager{AbstractDataManager: AbstractDataManager{TableModel{AbstractModel(Execution{})}}}
}

func (execut ExecutionEntityManagerImpl) GetDataManager() DataManagers {
	return executionDataManager
}
func (execut ExecutionEntityManagerImpl) FindById(entityId int64) ExecutionEntityImpl {
	execution := &ProcessInstance{}
	execut.GetDataManager().FindById(entityId, execution)
	entityImpl := ExecutionEntityImpl{}
	entityImpl.SetId(execution.Id)
	entityImpl.SetProcessDefineId(execution.ProcessDefineId)
	return entityImpl
}
