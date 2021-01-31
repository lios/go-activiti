package interceptor

import (
	"github.com/lios/go-activiti/engine/impl/invocation/agenda"
)

type CommandContextFactory struct {
}

func (factory CommandContextFactory) CreateCommandContext(command Command) CommandContext {
	context := CommandContext{Command: command, Agenda: &agenda.DefaultActivitiEngineAgenda{}}
	return context
}
