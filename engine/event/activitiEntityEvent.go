package event

type ActivitiEntityEvent interface {
	ActivitiEvent
	GetEntity() interface{}
}
