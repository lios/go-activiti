package converter

import (
	. "encoding/xml"
	"github.com/lios/go-activiti/engine/contanst"
	. "github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type ExclusiveGatewayXMLConverter struct {
	BpmnXMLConverter
}

func (exclusiveGateway ExclusiveGatewayXMLConverter) GetXMLElementName() string {
	return contanst.ELEMENT_GATEWAY_EXCLUSIVE
}
func (exclusiveGateway ExclusiveGatewayXMLConverter) ConvertXMLToElement(decoder *Decoder, token StartElement, model *BpmnModel, activeProcess *Process) BaseElement {
	exclusive := ExclusiveGateway{Gateway{FlowNode: FlowNode{BaseHandlerType: BaseHandlerType(ExclusiveGateway{}), IncomingFlow: make([]FlowElement, 0), OutgoingFlow: make([]FlowElement, 0)}}}
	decoder.DecodeElement(&exclusive, &token)
	return &exclusive
}
