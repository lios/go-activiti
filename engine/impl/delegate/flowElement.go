package delegate

//接口
type FlowElement interface {
	BaseElement
	SetOutgoing(f []FlowElement)
	SetIncoming(f []FlowElement)
	GetIncoming() []FlowElement
	GetOutgoing() []FlowElement
	GetBehavior() ActivityBehavior
	SetBehavior(behavior ActivityBehavior)

	SetSourceFlowElement(f FlowElement)
	SetTargetFlowElement(f FlowElement)
	GetSourceFlowElement() FlowElement
	GetTargetFlowElement() FlowElement
	GetHandlerType() string
}
