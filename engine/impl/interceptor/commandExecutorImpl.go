package interceptor

var commandExecutorImpl CommandExecutorImpl

type CommandExecutorImpl struct {
	First CommandInterceptor
}

func SetCommandExecutorImpl(commandExecutor CommandExecutorImpl) {
	commandExecutorImpl = commandExecutor
}

func GetCommandExecutorImpl() CommandExecutorImpl {
	return commandExecutorImpl
}

func (comm CommandExecutorImpl) Exe(conf Command) (interface{}, error) {
	return comm.First.Execute(conf)
}
