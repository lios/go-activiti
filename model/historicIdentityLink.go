package model

type HistoricIdentityLink struct {
	Id                int64
	Type              string `gorm:"column:type"`
	TaskId            int64  `gorm:"column:task_id"`
	ProcessInstanceId int    `gorm:"column:proc_inst_id"`
	GroupId           string `gorm:"column:group_id"`
	UserId            string `gorm:"column:user_id"`
}

func (HistoricIdentityLink) TableName() string {
	return "hi_identity_link"
}
