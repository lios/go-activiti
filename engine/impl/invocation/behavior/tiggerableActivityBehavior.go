package behavior

import (
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type TriggerableActivityBehavior interface {
	Trigger(execution delegate.DelegateExecution)
}
