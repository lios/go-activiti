package bpmn

//接口
type FlowElement interface {
	BaseElement
	SetOutgoing(f []FlowElement)
	SetIncoming(f []FlowElement)
	GetIncoming() []FlowElement
	GetOutgoing() []FlowElement

	SetSourceFlowElement(f FlowElement)
	SetTargetFlowElement(f FlowElement)
	GetSourceFlowElement() FlowElement
	GetTargetFlowElement() FlowElement
}
