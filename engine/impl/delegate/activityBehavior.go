package delegate

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type ActivityBehavior interface {
	Execute(execution entity.ExecutionEntity) error
}
