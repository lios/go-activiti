package model

import (
	"time"
)

type HistoricTask struct {
	Id                int64
	TaskId            int64     `gorm:"column:task_id"`
	TaskDefineKey     string    `gorm:"column:task_define_key"`
	TaskDefineName    string    `gorm:"column:task_define_name"`
	TenantId          string    `gorm:"column:tenant_id"`
	DeploymentId      int       `gorm:"column:deployment_id"`
	StartTime         time.Time `gorm:"column:start_time"`
	EndTime           time.Time `gorm:"column:end_time","default: null"`
	Assignee          string    `gorm:"column:assignee"`
	ProcessInstanceId int64     `gorm:"column:proc_inst_id"`
}

func (HistoricTask) TableName() string {
	return "hi_task"
}

func (HistoricTask) getTableName() string {
	return "hi_task"
}
