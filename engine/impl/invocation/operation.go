package invocation

type Operation interface {
	Run() error
}
