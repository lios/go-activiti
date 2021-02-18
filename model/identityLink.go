package model

type IdentityLink struct {
	Id                int64
	Type              string `gorm:"column:type"`
	TaskId            int64  `gorm:"column:task_id"`
	ProcessInstanceId int64  `gorm:"column:proc_inst_id"`
	GroupId           string `gorm:"column:group_id"`
	UserId            string `gorm:"column:user_id"`
}

func (IdentityLink) TableName() string {
	return "identity_link"
}

func (IdentityLink) getTableName() string {
	return "identity_link"
}
