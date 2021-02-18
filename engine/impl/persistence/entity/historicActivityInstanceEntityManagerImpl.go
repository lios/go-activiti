package entity

import (
	"github.com/lios/go-activiti/engine/impl/delegate"
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	. "github.com/lios/go-activiti/model"
	"time"
)

var (
	historicActivityInstanceEntity HistoricActivityInstanceEntity
	historicActinstDataManager     HistoricActinstDataManager
)

type HistoricActivityInstanceEntityManagerImpl struct {
	DefaultHistoryManager
}

func init() {
	//historicActinstDataManager = HistoricActinstDataManager{}
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

func (historicActivityInstanceEntityManager HistoricActivityInstanceEntityManagerImpl) RecordActivityStart(entity ExecutionEntity) {
	manager := historicActivityInstanceEntityManager.GetDataManager()
	actinstDataManager := manager.(HistoricActinstDataManager)
	historicActinst := HistoricActinst{}
	historicActinst.ProcessDefineId = entity.GetProcessDefineId()
	historicActinst.ProcessInstanceId = entity.GetProcessInstanceId()
	historicActinst.TaskId = entity.GetTaskId()
	historicActinst.ActId = entity.GetCurrentActivityId()
	historicActinst.StartTime = time.Now()
	if entity.GetCurrentFlowElement() != nil {
		historicActinst.ActName = entity.GetCurrentFlowElement().GetName()
		historicActinst.ActType = parseActivityType(entity.GetCurrentFlowElement())
	}
	actinstDataManager.Insert(&historicActinst)

}

func (historicActivityInstanceEntityManager HistoricActivityInstanceEntityManagerImpl) RecordTaskCreated(element delegate.FlowElement, entity ExecutionEntity) (err error) {
	var actinst = HistoricActinst{}
	manager := historicActivityInstanceEntityManager.GetDataManager()
	actinstDataManager := manager.(HistoricActinstDataManager)
	actinst, err = actinstDataManager.FindUnfinishedHistoricActivityInstancesByExecutionAndActivityId(entity.GetProcessInstanceId(), element.GetId())
	if err == nil {
		actinst.EndTime = time.Now()
		err = actinstDataManager.Update(actinst)
	}
	return err
}

func parseActivityType(element delegate.FlowElement) string {
	return element.GetHandlerType()
}

func (historicActivityInstanceEntityManager HistoricActivityInstanceEntityManagerImpl) RecordTaskId(task Task) {
	manager := historicActivityInstanceEntityManager.GetDataManager()
	actinstDataManager := manager.(HistoricActinstDataManager)
	actinst, err := actinstDataManager.FindUnfinishedHistoricActivityInstancesByExecutionAndActivityId(task.ProcessInstanceId, task.TaskDefineKey)
	if err == nil {
		actinst.Assignee = task.Assignee
		actinst.TaskId = task.Id
		err = actinstDataManager.Update(actinst)
	}
}
