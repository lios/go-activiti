package agenda

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type AbstractOperation struct {
	Execution entity.ExecutionEntity
}
