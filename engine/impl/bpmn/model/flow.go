package model

//父类实现体
type Flow struct {
	BaseElement
	Id                string `xml:"id,attr"`
	Name              string `xml:"name,attr"`
	IncomingFlow      []FlowElement
	OutgoingFlow      []FlowElement
	SourceFlowElement FlowElement
	TargetFlowElement FlowElement
}

func (flow *Flow) SetIncoming(f []FlowElement) {
	flow.IncomingFlow = f
}
func (flow *Flow) SetOutgoing(f []FlowElement) {
	flow.OutgoingFlow = f
}

func (flow *Flow) GetIncoming() []FlowElement {
	return flow.IncomingFlow
}
func (flow *Flow) GetOutgoing() []FlowElement {
	return flow.OutgoingFlow
}

func (flow *Flow) SetSourceFlowElement(f FlowElement) {
	flow.SourceFlowElement = f
}
func (flow *Flow) SetTargetFlowElement(f FlowElement) {
	flow.TargetFlowElement = f
}

func (flow *Flow) GetSourceFlowElement() FlowElement {
	return flow.SourceFlowElement
}
func (flow *Flow) GetTargetFlowElement() FlowElement {
	return flow.TargetFlowElement
}
func (flow *Flow) GetId() string {
	return flow.Id
}

func (flow *Flow) GetName() string {
	return flow.Name
}
