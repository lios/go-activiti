package model

import (
	"time"
)

//流程实例
type HistoricActinst struct {
	TableModel
	Id                int64
	ProcessDefineId   int64     `gorm:"column:process_define_id"`
	ProcessInstanceId int64     `gorm:"column:proc_inst_id"`
	TaskId            int64     `gorm:"column:task_id"`
	ActId             string    `gorm:"column:act_id"`
	ActName           string    `gorm:"column:act_name"`
	ActType           string    `gorm:"column:act_type"`
	TenantId          string    `gorm:"column:tenant_id"`
	StartTime         time.Time `gorm:"column:start_time"`
	EndTime           time.Time `gorm:"column:end_time","default: null"`
	StartUserId       string    `gorm:"column:start_user_id"`
	Assignee          string    `gorm:"column:assignee"`
}

func (HistoricActinst) TableName() string {
	return "hi_actinst"
}

func (HistoricActinst) getTableName() string {
	return "hi_actinst"
}
