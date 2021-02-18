package model

import "time"

type Execution struct {
	Id                int64
	ProcessInstanceId int64     `gorm:"column:proc_inst_id"`
	EndTime           time.Time `gorm:"column:end_time","default: null"`
}

func (Execution) TableName() string {
	return "execution"
}

func (Execution) getTableName() string {
	return "execution"
}
