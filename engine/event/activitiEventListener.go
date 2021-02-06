package event

type ActivitiEventListener interface {
	OnEvent(event ActivitiEvent) error
}
