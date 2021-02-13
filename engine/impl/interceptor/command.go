package interceptor

type Command interface {
	Execute(interceptor CommandContext) (interface{}, error)
}
