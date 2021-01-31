package engine

type ProcessEngine interface {
	GetRuntimeService() RuntimeService

	GetTaskService() TaskService
}
