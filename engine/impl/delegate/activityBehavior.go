package delegate

type ActivityBehavior interface {
	Execute(execution DelegateExecution) error
}
