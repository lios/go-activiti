package behavior

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type FlowNodeActivityBehavior interface {
	Leave(execution entity.ExecutionEntity) error
}
