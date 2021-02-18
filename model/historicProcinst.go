package model

import (
	"time"
)

//流程实例
type HistoricProcess struct {
	Id                int64
	ProcessInstanceId int64     `gorm:"column:proc_inst_id"`
	Key               string    `gorm:"column:key"`
	Name              string    `gorm:"column:name"`
	BusinessKey       string    `gorm:"column:business_key"`
	TenantId          string    `gorm:"column:tenant_id"`
	DeploymentId      int64     `gorm:"column:deployment_id"`
	StartTime         time.Time `gorm:"column:start_time"`
	EndTime           time.Time `gorm:"column:end_time","default: null"`
	StartUserId       string    `gorm:"column:start_user_id"`
	ProcessDefineId   int64     `gorm:"column:process_define_id"`
}

func (HistoricProcess) TableName() string {
	return "hi_process_instance"
}

func (HistoricProcess) getTableName() string {
	return "hi_process_instance"
}
