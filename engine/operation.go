package engine

type Operation interface {
	Run() error
}
