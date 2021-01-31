package interceptor

import (
	"github.com/lios/go-activiti/engine"
	"github.com/lios/go-activiti/engine/impl/invocation"
)

type CommandContextInterceptor struct {
	Next                       CommandInterceptor
	ProcessEngineConfiguration *engine.ProcessEngineConfiguration
	CommandContextFactory      CommandContextFactory
}

func (commandContext CommandContextInterceptor) Execute(command Command) (interface{}, error) {
	cContext, err := GetCommandContext()
	if err != nil {
		cContext = commandContext.CommandContextFactory.CreateCommandContext(command)
	}
	invocation.SetAgenda(cContext.Agenda)
	SetCommandContext(cContext)
	return commandContext.Next.Execute(command)
}

func (commandContext *CommandContextInterceptor) SetNext(next CommandInterceptor) {
	commandContext.Next = next
}
