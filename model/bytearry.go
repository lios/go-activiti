package model

//流程源数据
type Bytearry struct {
	Id           int64  `gorm:"column:id"`
	Key          string `gorm:"column:key"`
	Name         string `gorm:"column:name"`
	Version      int    `gorm:"column:version"`
	Bytes        string `gorm:"column:bytes"`
	DeploymentId int64  `gorm:"column:deployment_id"`
}

func (Bytearry) TableName() string {
	return "bytearry"
}
