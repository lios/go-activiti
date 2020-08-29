package behavior

type CommandExecutor interface {
	Exe(conf Command) (interface{}, error)
}
