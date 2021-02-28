package entity

import "time"

type TaskEntity interface {
	Entity
	GetTaskName() string

	SetAssignee(assignee string)

	SetStartTime(startTime time.Time)

	SetTaskDefineKey(taskDefineKey string)

	GetTaskDefineKey() string

	SetTaskDefineName(taskDefineName string)
}
