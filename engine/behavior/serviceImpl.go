package behavior

var serviceImpl ServiceImpl

type ServiceImpl struct {
	CommandExecutor CommandExecutor
}

func (serviceImpl *ServiceImpl) SetCommandExecutor(commandExecutor CommandExecutor) {
	serviceImpl.CommandExecutor = commandExecutor
}

func GetServiceImpl() ServiceImpl {
	return serviceImpl
}

func SetServiceImpl(service ServiceImpl) {
	serviceImpl = service
}
