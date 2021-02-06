package entity

type HistoricTaskInstanceEntityImpl struct {
	AbstractEntity
	Assignee string
}

func (historicTaskInstanceEntity HistoricTaskInstanceEntityImpl) SetAssignee(assignee string) {
	historicTaskInstanceEntity.Assignee = assignee
}
