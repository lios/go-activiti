package converter

import (
	. "encoding/xml"
	"github.com/lios/go-activiti/engine/contanst"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type InclusiveGatewayXMLConverter struct {
	BpmnXMLConverter
}

func (inclusiveGateway InclusiveGatewayXMLConverter) GetXMLElementName() string {
	return contanst.ELEMENT_GATEWAY_INCLUSIVE
}
func (inclusiveGateway InclusiveGatewayXMLConverter) ConvertXMLToElement(decoder *Decoder, token StartElement, model *BpmnModel, activeProcess *Process) delegate.BaseElement {
	inclusive := InclusiveGateway{Gateway{FlowNode: FlowNode{BaseHandlerType: delegate.BaseHandlerType(InclusiveGateway{}), IncomingFlow: make([]delegate.FlowElement, 0), OutgoingFlow: make([]delegate.FlowElement, 0)}}}
	decoder.DecodeElement(&inclusive, &token)
	return &inclusive
}
