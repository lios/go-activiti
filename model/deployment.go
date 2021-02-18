package model

import "time"

//部署
type Deployment struct {
	Id         int64     `gorm:"column:id"`
	Key        string    `gorm:"column:key"`
	Name       string    `gorm:"column:name"`
	Version    int       `gorm:"column:version"`
	TenantId   string    `gorm:"column:tenant_id"`
	DeployTime time.Time `gorm:"column:deploy_time"`
}

func (Deployment) TableName() string {
	return "deployment"
}

func (Deployment) getTableName() string {
	return "deployment"
}
