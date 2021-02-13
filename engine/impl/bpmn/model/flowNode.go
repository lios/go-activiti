package model

import (
	. "github.com/lios/go-activiti/engine/impl/bpmn"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

//父类实现体
type FlowNode struct {
	BaseHandlerType
	Id                string `xml:"id,attr"`
	Name              string `xml:"name,attr"`
	IncomingFlow      []FlowElement
	OutgoingFlow      []FlowElement
	SourceFlowElement FlowElement
	TargetFlowElement FlowElement
	Behavior          delegate.ActivityBehavior
}

func (flow *FlowNode) SetIncoming(f []FlowElement) {
	flow.IncomingFlow = f
}
func (flow *FlowNode) SetOutgoing(f []FlowElement) {
	flow.OutgoingFlow = f
}

func (flow *FlowNode) GetIncoming() []FlowElement {
	return flow.IncomingFlow
}
func (flow *FlowNode) GetOutgoing() []FlowElement {
	return flow.OutgoingFlow
}

func (flow *FlowNode) SetSourceFlowElement(f FlowElement) {
	flow.SourceFlowElement = f
}
func (flow *FlowNode) SetTargetFlowElement(f FlowElement) {
	flow.TargetFlowElement = f
}

func (flow *FlowNode) GetSourceFlowElement() FlowElement {
	return flow.SourceFlowElement
}
func (flow *FlowNode) GetTargetFlowElement() FlowElement {
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
