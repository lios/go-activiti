package converter

import (
	. "encoding/xml"
	"github.com/lios/go-activiti/engine/contanst"
	. "github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type ParallelGatewayXMLConverter struct {
	BpmnXMLConverter
}

func (parallelGateway ParallelGatewayXMLConverter) GetXMLElementName() string {
	return contanst.ELEMENT_GATEWAY_PARALLEL
}
func (parallelGateway ParallelGatewayXMLConverter) ConvertXMLToElement(decoder *Decoder, token StartElement, model *BpmnModel, activeProcess *Process) BaseElement {
	parallel := ParallelGateway{Gateway{FlowNode: FlowNode{BaseHandlerType: BaseHandlerType(ParallelGateway{}), IncomingFlow: make([]FlowElement, 0), OutgoingFlow: make([]FlowElement, 0)}}}
	decoder.DecodeElement(&parallel, &token)
	return &parallel
}
