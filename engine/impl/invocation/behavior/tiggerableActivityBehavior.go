package behavior

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type TriggerableActivityBehavior interface {
	Trigger(entity entity.ExecutionEntity)
}
