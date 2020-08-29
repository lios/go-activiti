package event

import . "github.com/lios/go-activiti/errs"

type ActivitiEventSupport struct {
	EventListeners []ActivitiEventListener
}

func (activitiEventSupport *ActivitiEventSupport) AddEventListener(listenerToAdd ActivitiEventListener) (err error) {
	if listenerToAdd == nil {
		err = ProcessError{Msg: "Listener cannot be null."}
	}
	activitiEventSupport.EventListeners = append(activitiEventSupport.EventListeners, listenerToAdd)
	return err
}

func (activitiEventSupport ActivitiEventSupport) DispatchEvent(event ActivitiEvent) (err error) {
	if event == nil {
		err = ProcessError{Msg: "Event cannot be null."}
		return err
	}

	if len(event.GetType()) == 0 {
		err = ProcessError{Msg: "Event type cannot be null."}
		return err
	}

	// Call global listeners
	if activitiEventSupport.EventListeners != nil && len(activitiEventSupport.EventListeners) > 0 {
		for _, listener := range activitiEventSupport.EventListeners {
			err = dispatchEvent(event, listener)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func dispatchEvent(event ActivitiEvent, listener ActivitiEventListener) error {
	return listener.OnEvent(event)
}
