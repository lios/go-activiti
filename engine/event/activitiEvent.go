package event

type ActivitiEventType string

const (
	TASK_CREATED ActivitiEventType = "TASK_CREATED"

	TASK_ASSIGNED ActivitiEventType = "TASK_ASSIGNED"

	TASK_COMPLETED ActivitiEventType = "TASK_COMPLETED"
)

type ActivitiEvent interface {
	GetType() ActivitiEventType
}
