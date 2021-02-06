package entity

import . "github.com/lios/go-activiti/engine/impl/persistence/entity/data"

var (
	historicActivityInstanceEntity HistoricActivityInstanceEntity
	historicActinstDataManager     HistoricActinstDataManager
)

type HistoricActivityInstanceEntityManagerImpl struct {
	DefaultHistoryManager
}

func (historicActivityInstanceEntityManager HistoricActivityInstanceEntityManagerImpl) GetDataManager() DataManagers {
	return historicActinstDataManager
}

func (historicActivityInstanceEntityManager HistoricActivityInstanceEntityManagerImpl) DeleteHistoricActivityInstancesByProcessInstanceId(historicProcessInstanceId string) {

}

func (historicActivityInstanceEntityManager HistoricActivityInstanceEntityManagerImpl) RecordEnd(taskId int64) {
	manager := historicActivityInstanceEntityManager.GetDataManager()
	actinstDataManager := manager.(HistoricActinstDataManager)
	actinstDataManager.UpdateTaskId(taskId)
}
