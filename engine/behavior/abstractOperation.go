package behavior

import (
	"github.com/lios/go-activiti/engine"
)

type AbstractOperation struct {
	CommandContext CommandContext
	Execution      engine.ExecutionEntity
}
