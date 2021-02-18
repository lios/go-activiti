package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	. "github.com/lios/go-activiti/model"
)

var (
	historicTaskInstanceEntity HistoricTaskInstanceEntity
	historicTaskDataManager    HistoricTaskDataManager
)

func init() {
	historicTaskDataManager = HistoricTaskDataManager{AbstractDataManager: AbstractDataManager{TableModel{AbstractModel(HistoricTask{})}}}
}

type HistoricTaskInstanceEntityManagerImpl struct {
	DefaultHistoryManager
	AbstractEntityManager
}

func (historicTaskInstanceEntityManager HistoricTaskInstanceEntityManagerImpl) GetDataManager() DataManagers {
	return historicTaskDataManager
}
func (historicTaskInstanceEntityManager HistoricTaskInstanceEntityManagerImpl) delete(taskId string) {

}
