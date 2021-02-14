package entity

type HistoricActivityInstanceEntityManager interface {
	DeleteHistoricActivityInstancesByProcessInstanceId(historicProcessInstanceId string)

	RecordEnd(taskId int64)

	RecordActivityStart(entity ExecutionEntity)
}
