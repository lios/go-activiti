package entity

type HistoryManager interface {
	RecordTaskEnd(taskId int64, deleteReason string)
}
