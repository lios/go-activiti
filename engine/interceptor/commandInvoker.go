package interceptor

type CommandInvoker struct {
	Next CommandInterceptor
}

func (commandInvoker CommandInvoker) Execute(command Command) (result interface{}, err error) {
	context, err := GetCommandContext()
	if err != nil {
		return nil, err
	}
	result, err = command.Execute(context)
	if err != nil {
		return nil, err
	}
	err = executeOperations(context)
	return result, err
}

func executeOperations(context CommandContext) (err error) {
	for !context.Agenda.IsEmpty() {
		err = context.Agenda.GetNextOperation().Run()
		if err != nil {
			return err
		}
	}
	return err
}

func (commandInvoker *CommandInvoker) SetNext(next CommandInterceptor) {
	commandInvoker.Next = next
}
