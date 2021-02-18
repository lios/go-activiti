package model

import (
	"github.com/lios/go-activiti/engine/impl/delegate"
)

//父类实现体
type FlowNode struct {
	delegate.BaseHandlerType
	Id                string `xml:"id,attr"`
	Name              string `xml:"name,attr"`
	IncomingFlow      []delegate.FlowElement
	OutgoingFlow      []delegate.FlowElement
	SourceFlowElement delegate.FlowElement
	TargetFlowElement delegate.FlowElement
	Behavior          delegate.ActivityBehavior
}

func (flow *FlowNode) SetIncoming(f []delegate.FlowElement) {
	flow.IncomingFlow = f
}
func (flow *FlowNode) SetOutgoing(f []delegate.FlowElement) {
	flow.OutgoingFlow = f
}

func (flow *FlowNode) GetIncoming() []delegate.FlowElement {
	return flow.IncomingFlow
}
func (flow *FlowNode) GetOutgoing() []delegate.FlowElement {
	return flow.OutgoingFlow
}

func (flow *FlowNode) SetSourceFlowElement(f delegate.FlowElement) {
	flow.SourceFlowElement = f
}
func (flow *FlowNode) SetTargetFlowElement(f delegate.FlowElement) {
	flow.TargetFlowElement = f
}

func (flow *FlowNode) GetSourceFlowElement() delegate.FlowElement {
	return flow.SourceFlowElement
}
func (flow *FlowNode) GetTargetFlowElement() delegate.FlowElement {
	return flow.TargetFlowElement
}

func (flow *FlowNode) GetId() string {
	return flow.Id
}

func (flow *FlowNode) GetName() string {
	return flow.Name
}

func (flow *FlowNode) GetBehavior() delegate.ActivityBehavior {
	return flow.Behavior
}

func (flow *FlowNode) SetBehavior(behavior delegate.ActivityBehavior) {
	flow.Behavior = behavior
}

func (flow *FlowNode) GetHandlerType() string {
	return flow.BaseHandlerType.GetType()
}
