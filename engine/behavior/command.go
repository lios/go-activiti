package behavior

type Command interface {
	Execute(interceptor CommandContext) (interface{}, error)
}
