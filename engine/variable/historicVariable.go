package variable

import (
	"time"
)

type HistoricVariable struct {
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

func (HistoricVariable) TableName() string {
	return "hi_variable"
}

func (variable HistoricVariable) GetName() string {
	return variable.Name
}

func (variable HistoricVariable) GetProcessInstanceId() int64 {
	return variable.ProcessInstanceId
}

func (variable HistoricVariable) GetTaskId() int64 {
	return variable.TaskId
}

func (variable HistoricVariable) GetNumberValue() int {
	return variable.Number
}

func (variable *HistoricVariable) SetNumberValue(value int) {
	variable.Number = value
}

func (variable HistoricVariable) GetTextValue() string {
	return variable.Text
}

func (variable *HistoricVariable) SetTextValue(value string) {
	variable.Text = value
}

func (variable *HistoricVariable) SetValue(value interface{}, variableType VariableType) {
	variableType.SetValue(value, variable)
}

func (variable *HistoricVariable) SetBlobValue(value string) {
	variable.Blob = value
}

func (variable HistoricVariable) GetBlobValue() string {
	return variable.Blob
}
