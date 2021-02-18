package entity

import (
	"github.com/lios/go-activiti/engine/impl/delegate"
	. "github.com/lios/go-activiti/model"
)

type HistoricActivityInstanceEntityManager interface {
	DeleteHistoricActivityInstancesByProcessInstanceId(historicProcessInstanceId string)

	RecordEnd(taskId int64)

	RecordActivityStart(entity ExecutionEntity)

	RecordTaskCreated(element delegate.FlowElement, entity ExecutionEntity) (err error)

	RecordTaskId(task Task)
}
