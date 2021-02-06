package entity

import . "github.com/lios/go-activiti/engine/impl/persistence/entity/data"

type DefaultHistoryManager struct {
	AbstractEntityManager
}

func (defaultHistoryManager DefaultHistoryManager) RecordTaskEnd(taskId int64, deleteReason string) {
	manager := GetHistoricTaskInstanceEntityManager()
	dataManager := manager.GetDataManager()
	historicTaskDataManager := dataManager.(HistoricTaskDataManager)
	historicTask, err := historicTaskDataManager.GetByTaskId(taskId)
	if err != nil {
		return
	}
	historicTaskDataManager.MarkEnded(historicTask.Id)
}
