package interceptor

type CommandExecutor interface {
	Exe(conf Command) (interface{}, error)
}
