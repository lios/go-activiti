package entity

type HistoricActivityInstanceEntityImpl struct {
	AbstractEntity
	ActivityId string
}

func (historicActivityInstanceEntity HistoricActivityInstanceEntityImpl) SetActivityId(activityId string) {
	historicActivityInstanceEntity.ActivityId = activityId
}
