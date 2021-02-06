package cfg

import "github.com/lios/go-activiti/engine/interceptor"

var commandExecutorImpl CommandExecutorImpl

type CommandExecutorImpl struct {
	First interceptor.CommandInterceptor
}

func SetCommandExecutorImpl(commandExecutor CommandExecutorImpl) {
	commandExecutorImpl = commandExecutor
}

func GetCommandExecutorImpl() CommandExecutorImpl {
	return commandExecutorImpl
}

func (comm CommandExecutorImpl) Exe(conf interceptor.Command) (interface{}, error) {
	return comm.First.Execute(conf)
}
