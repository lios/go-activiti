package model

import (
	"time"
)

type Task struct {
	TableModel
	Id                int64
	TaskDefineKey     string    `gorm:"column:task_define_key"`
	TaskDefineName    string    `gorm:"column:task_define_name"`
	TenantId          string    `gorm:"column:tenant_id"`
	DeploymentId      int       `gorm:"column:deployment_id"`
	StartTime         time.Time `gorm:"column:start_time"`
	Assignee          string    `gorm:"column:assignee"`
	ProcessInstanceId int64     `gorm:"column:proc_inst_id"`
}

func (Task) TableName() string {
	return "ru_task"
}

func (Task) getTableName() string {
	return "ru_task"
}
