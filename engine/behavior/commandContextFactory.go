package behavior

type CommandContextFactory struct {
}

func (factory CommandContextFactory) CreateCommandContext(command Command) CommandContext {
	context := CommandContext{Command: command, Agenda: &DefaultActivitiEngineAgenda{}}
	return context
}
