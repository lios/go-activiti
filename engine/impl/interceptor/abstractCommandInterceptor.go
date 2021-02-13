package interceptor

type AbstractCommandInterceptor struct {
	Next CommandInterceptor
}
