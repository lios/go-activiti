package event

type ActivitiEventDispatcher interface {
	AddEventListener(listenerToAdd ActivitiEventListener)

	RemoveEventListener(listenerToRemove ActivitiEventListener)

	DispatchEvent(event ActivitiEvent)

	SetEnabled(enabled bool)

	IsEnabled() bool
}
