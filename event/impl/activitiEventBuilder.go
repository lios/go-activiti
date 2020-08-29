package impl

import (
	"github.com/lios/go-activiti/errs"
	. "github.com/lios/go-activiti/event"
)

type ActivitiEventBuilder struct {
}

func CreateEvent() ActivitiEvent {
	return nil
}

func CreateEntityEvent(eventType ActivitiEventType, entity interface{}) (ActivitiEntityEvent, error) {
	entityEventImpl := ActivitiEntityEventImpl{}
	entityEventImpl.ActivitiEventImpl = ActivitiEventImpl{}
	entityEventImpl.EventType = eventType
	var err error = nil
	if entity == nil {
		err = errs.ProcessError{Msg: "Entity cannot be null."}
	}
	entityEventImpl.Entity = entity
	return entityEventImpl, err
}
