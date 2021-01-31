package cfg

import "github.com/lios/go-activiti/engine/interceptor"

var serviceImpl ServiceImpl

var CommandExecutor interceptor.CommandExecutor

type ServiceImpl struct {
}

func (serviceImpl *ServiceImpl) SetCommandExecutor(commandExecutor interceptor.CommandExecutor) {
	CommandExecutor = commandExecutor
}

func GetServiceImpl() ServiceImpl {
	return serviceImpl
}

func SetServiceImpl(service ServiceImpl) {
	serviceImpl = service
}

func (serviceImpl *ServiceImpl) GetCommandExecutor() interceptor.CommandExecutor {
	return CommandExecutor
}
