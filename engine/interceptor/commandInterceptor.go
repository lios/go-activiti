package interceptor

type CommandInterceptor interface {
	Execute(command Command) (interface{}, error)

	SetNext(next CommandInterceptor)
}
