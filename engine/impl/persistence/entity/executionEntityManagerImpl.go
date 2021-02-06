package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	"github.com/lios/go-activiti/model"
)

var executionDataManager ExecutionDataManager

type ExecutionEntityManagerImpl struct {
	AbstractEntityManager
}

func (execut ExecutionEntityManagerImpl) GetDataManager() DataManagers {
	return executionDataManager
}
func (execut ExecutionEntityManagerImpl) FindById(entityId int64) ExecutionEntityImpl {
	execution := &model.Execution{}
	execut.GetDataManager().FindById(entityId, execution)
	entityImpl := ExecutionEntityImpl{}
	entityImpl.SetId(execution.Id)
	return entityImpl
}
