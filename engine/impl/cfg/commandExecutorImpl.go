package cfg

import "github.com/lios/go-activiti/engine/interceptor"

type CommandExecutorImpl struct {
	First interceptor.CommandInterceptor
}

func (comm CommandExecutorImpl) Exe(conf interceptor.Command) (interface{}, error) {
	return comm.First.Execute(conf)
}
