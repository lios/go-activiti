package delegate

import "github.com/lios/go-activiti/engine/delegate"

type ActivityBehavior interface {
	Execute(execution delegate.DelegateExecution) error
}
