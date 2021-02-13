package entity

import "time"

type TaskEntity interface {
	Entity
	GetTaskName() string

	SetAssignee(assignee string)

	SetStartTime(startTime time.Time)

	SetTaskDefineKey(taskDefineKey string)

	SetTaskDefineName(taskDefineName string)

	SetExecutionVariables(parameters map[string]interface{}) error
}
