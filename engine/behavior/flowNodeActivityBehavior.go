package behavior

import (
	"github.com/lios/go-activiti/engine"
)

type FlowNodeActivityBehavior interface {
	Leave(execution engine.ExecutionEntity) error
}
