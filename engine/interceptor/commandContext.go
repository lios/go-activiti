package interceptor

import (
	"github.com/lios/go-activiti/engine/impl/invocation"
)

type CommandContext struct {
	Command Command
	Agenda  invocation.ActivitiEngineAgenda
}
