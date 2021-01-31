package engine

type ProcessEngineConfiguration interface {
	GetRuntimeService() RuntimeService
}
