package behavior

import (
	"github.com/lios/go-activiti/engine"
)

type TriggerableActivityBehavior interface {
	Trigger(entity engine.ExecutionEntity)
}
