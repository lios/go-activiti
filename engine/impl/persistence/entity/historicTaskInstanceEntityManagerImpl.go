package entity

import . "github.com/lios/go-activiti/engine/impl/persistence/entity/data"

var (
	historicTaskInstanceEntity HistoricTaskInstanceEntity
	historicTaskDataManager    HistoricTaskDataManager
)

type HistoricTaskInstanceEntityManagerImpl struct {
	DefaultHistoryManager
}

func (historicTaskInstanceEntityManager HistoricTaskInstanceEntityManagerImpl) GetDataManager() DataManagers {
	return historicTaskDataManager
}
func (historicTaskInstanceEntityManager HistoricTaskInstanceEntityManagerImpl) delete(taskId string) {

}
