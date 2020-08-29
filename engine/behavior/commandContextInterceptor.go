package behavior

type CommandContextInterceptor struct {
	Next                       CommandInterceptor
	ProcessEngineConfiguration *ProcessEngineConfiguration
	CommandContextFactory      CommandContextFactory
}

func (commandContext CommandContextInterceptor) Execute(command Command) (interface{}, error) {
	context, err := GetCommandContext()
	if err != nil {
		context = commandContext.CommandContextFactory.CreateCommandContext(command)
	}
	SetCommandContext(context)
	return commandContext.Next.Execute(command)
}

func (commandContext *CommandContextInterceptor) SetNext(next CommandInterceptor) {
	commandContext.Next = next
}
