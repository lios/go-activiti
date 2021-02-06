package event

var activitiEventManager ActivitiEventManager

type ActivitiEventManager struct {
	EventDispatcher ActivitiEventDispatcher
	EventListeners  []ActivitiEventListener
}

func init() {
	activitiEventManager = ActivitiEventManager{}
}

func SetAtivitiEventDispatcher(eventDispatcher ActivitiEventDispatcher) {
	activitiEventManager.EventDispatcher = eventDispatcher
}

func SetEventListeners(eventListeners []ActivitiEventListener) {
	activitiEventManager.EventListeners = eventListeners
}

func GetAtivitiEventDispatcher() ActivitiEventDispatcher {
	return activitiEventManager.EventDispatcher
}

func GetEventListeners() []ActivitiEventListener {
	return activitiEventManager.EventListeners
}
