package converter

import (
	. "encoding/xml"
	"github.com/lios/go-activiti/engine/contanst"
	. "github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type InclusiveGatewayXMLConverter struct {
	BpmnXMLConverter
}

func (inclusiveGateway InclusiveGatewayXMLConverter) GetXMLElementName() string {
	return contanst.ELEMENT_GATEWAY_INCLUSIVE
}
func (inclusiveGateway InclusiveGatewayXMLConverter) ConvertXMLToElement(decoder *Decoder, token StartElement, model *BpmnModel, activeProcess *Process) BaseElement {
	inclusive := InclusiveGateway{Gateway{FlowNode: FlowNode{BaseHandlerType: BaseHandlerType(InclusiveGateway{}), IncomingFlow: make([]FlowElement, 0), OutgoingFlow: make([]FlowElement, 0)}}}
	decoder.DecodeElement(&inclusive, &token)
	return &inclusive
}
