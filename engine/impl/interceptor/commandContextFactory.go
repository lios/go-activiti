package interceptor

import (
	. "github.com/lios/go-activiti/engine/impl/invocation"
)

var activitiEngineAgenda ActivitiEngineAgenda

type CommandContextFactory struct {
}

func SetActivitiEngineAgenda(agenda ActivitiEngineAgenda) {
	activitiEngineAgenda = agenda
}
func (factory CommandContextFactory) CreateCommandContext(command Command) CommandContext {
	context := CommandContext{Command: command, Agenda: activitiEngineAgenda}
	return context
}
