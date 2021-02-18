package converter

import (
	. "encoding/xml"
	"github.com/lios/go-activiti/engine/contanst"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type StartEventXMLConverter struct {
	BpmnXMLConverter
}

func (start StartEventXMLConverter) GetXMLElementName() string {
	return contanst.ELEMENT_EVENT_START
}
func (start StartEventXMLConverter) ConvertXMLToElement(decoder *Decoder, token StartElement, model *BpmnModel, activeProcess *Process) delegate.BaseElement {
	startEvent := StartEvent{FlowNode: FlowNode{BaseHandlerType: delegate.BaseHandlerType(StartEvent{}), IncomingFlow: make([]delegate.FlowElement, 0), OutgoingFlow: make([]delegate.FlowElement, 0)}}
	decoder.DecodeElement(&startEvent, &token)
	activeProcess.InitialFlowElement = &startEvent
	return &startEvent
}
