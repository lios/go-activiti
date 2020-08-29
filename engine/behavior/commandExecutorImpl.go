package behavior

type CommandExecutorImpl struct {
	First CommandInterceptor
}

func (comm CommandExecutorImpl) Exe(conf Command) (interface{}, error) {
	return comm.First.Execute(conf)
}
