package converter

import (
	. "encoding/xml"
	"github.com/lios/go-activiti/engine/contanst"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type ExclusiveGatewayXMLConverter struct {
	BpmnXMLConverter
}

func (exclusiveGateway ExclusiveGatewayXMLConverter) GetXMLElementName() string {
	return contanst.ELEMENT_GATEWAY_EXCLUSIVE
}
func (exclusiveGateway ExclusiveGatewayXMLConverter) ConvertXMLToElement(decoder *Decoder, token StartElement, model *BpmnModel, activeProcess *Process) delegate.BaseElement {
	exclusive := ExclusiveGateway{Gateway{FlowNode: FlowNode{BaseHandlerType: delegate.BaseHandlerType(ExclusiveGateway{}), IncomingFlow: make([]delegate.FlowElement, 0), OutgoingFlow: make([]delegate.FlowElement, 0)}}}
	decoder.DecodeElement(&exclusive, &token)
	return &exclusive
}
