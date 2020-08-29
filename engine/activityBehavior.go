package engine

type ActivityBehavior interface {
	Execute(execution ExecutionEntity) error
}
