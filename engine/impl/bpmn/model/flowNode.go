package model

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

//父类实现体
type FlowNode struct {
	Id                string `xml:"id,attr"`
	Name              string `xml:"name,attr"`
	IncomingFlow      []bpmn.FlowElement
	OutgoingFlow      []bpmn.FlowElement
	SourceFlowElement bpmn.FlowElement
	TargetFlowElement bpmn.FlowElement
	Behavior          delegate.ActivityBehavior
}

func (flow FlowNode) SetIncoming(f []bpmn.FlowElement) {
	flow.IncomingFlow = f
}
func (flow FlowNode) SetOutgoing(f []bpmn.FlowElement) {
	flow.OutgoingFlow = f
}

func (flow FlowNode) GetIncoming() []bpmn.FlowElement {
	return flow.IncomingFlow
}
func (flow FlowNode) GetOutgoing() []bpmn.FlowElement {
	return flow.OutgoingFlow
}

func (flow FlowNode) SetSourceFlowElement(f bpmn.FlowElement) {
	flow.SourceFlowElement = f
}
func (flow FlowNode) SetTargetFlowElement(f bpmn.FlowElement) {
	flow.TargetFlowElement = f
}

func (flow FlowNode) GetSourceFlowElement() bpmn.FlowElement {
	return flow.SourceFlowElement
}
func (flow FlowNode) GetTargetFlowElement() bpmn.FlowElement {
	return flow.TargetFlowElement
}

func (flow FlowNode) GetId() string {
	return flow.Id
}

func (flow FlowNode) GetName() string {
	return flow.Name
}

func (flow FlowNode) GetBehavior() delegate.ActivityBehavior {
	return flow.Behavior
}

func (flow FlowNode) SetBehavior(behavior delegate.ActivityBehavior) {
	flow.Behavior = behavior
}
