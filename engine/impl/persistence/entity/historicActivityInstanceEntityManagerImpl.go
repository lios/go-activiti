package entity

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	"github.com/lios/go-activiti/model"
	"reflect"
)

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

func (historicActivityInstanceEntityManager HistoricActivityInstanceEntityManagerImpl) RecordActivityStart(entity ExecutionEntity) {
	manager := historicActivityInstanceEntityManager.GetDataManager()
	actinstDataManager := manager.(HistoricActinstDataManager)
	historicActinst := model.HistoricActinst{}
	historicActinst.ProcessDefineId = entity.GetProcessDefineId()
	historicActinst.ProcessInstanceId = entity.GetProcessInstanceId()
	historicActinst.ActId = entity.GetCurrentActivityId()
	if entity.GetCurrentFlowElement() != nil {
		historicActinst.ActName = entity.GetCurrentFlowElement().GetName()
		historicActinst.ActType = parseActivityType(entity.GetCurrentFlowElement())
	}
	actinstDataManager.Insert(&historicActinst)

}

func parseActivityType(element bpmn.FlowElement) string {
	typeOf := reflect.TypeOf(element)
	return typeOf.Name()
}
