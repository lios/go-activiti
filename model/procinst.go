package model

import (
	"time"
)

//流程实例
type ProcessInstance struct {
	Id              int64
	Key             string    `gorm:"column:key"`
	Name            string    `gorm:"column:name"`
	BusinessKey     string    `gorm:"column:business_key"`
	TenantId        string    `gorm:"column:tenant_id"`
	DeploymentId    int64     `gorm:"column:deployment_id"`
	StartTime       time.Time `gorm:"column:start_time"`
	StartUserId     string    `gorm:"column:start_user_id"`
	ProcessDefineId int64     `gorm:"column:process_define_id"`
}

func (ProcessInstance) TableName() string {
	return "process_instance"
}

func (processInstance ProcessInstance) setBusinessKey(businessKey string) {
	processInstance.BusinessKey = businessKey
}
