package variable

import (
	"time"
)

type Variable struct {
	Id                int64
	Version           int64     `gorm:"column:version"`
	TaskId            int64     `gorm:"column:task_id"`
	ProcessInstanceId int64     `gorm:"column:proc_inst_id"`
	Name              string    `gorm:"column:name"`
	Type              string    `gorm:"column:type"`
	Date              time.Time `gorm:"column:date"`
	Number            int       `gorm:"column:number"`
	Float             float64   `gorm:"column:float"`
	Text              string    `gorm:"column:text"`
	Blob              string    `gorm:"column:blob"`
}

func (Variable) TableName() string {
	return "ru_variable"
}

func (Variable) getTableName() string {
	return "ru_variable"
}

func (variable Variable) GetName() string {
	return variable.Name
}

func (variable Variable) GetProcessInstanceId() int64 {
	return variable.ProcessInstanceId
}

func (variable Variable) GetTaskId() int64 {
	return variable.TaskId
}

func (variable Variable) GetNumberValue() int {
	return variable.Number
}

func (variable *Variable) SetNumberValue(value int) {
	variable.Number = value
}

func (variable Variable) GetTextValue() string {
	return variable.Text
}

func (variable *Variable) SetTextValue(value string) {
	variable.Text = value
}

func (variable *Variable) SetValue(value interface{}, variableType VariableType) {
	variableType.SetValue(value, variable)
}

func (variable *Variable) SetBlobValue(value string) {
	variable.Blob = value
}

func (variable Variable) GetBlobValue() string {
	return variable.Blob
}
