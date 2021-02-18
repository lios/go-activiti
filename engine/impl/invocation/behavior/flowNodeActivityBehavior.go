package behavior

import (
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type FlowNodeActivityBehavior interface {
	Leave(execution delegate.DelegateExecution) error
}
