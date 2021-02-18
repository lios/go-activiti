package converter

import (
	. "encoding/xml"
	"github.com/lios/go-activiti/engine/contanst"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type ParallelGatewayXMLConverter struct {
	BpmnXMLConverter
}

func (parallelGateway ParallelGatewayXMLConverter) GetXMLElementName() string {
	return contanst.ELEMENT_GATEWAY_PARALLEL
}
func (parallelGateway ParallelGatewayXMLConverter) ConvertXMLToElement(decoder *Decoder, token StartElement, model *BpmnModel, activeProcess *Process) delegate.BaseElement {
	parallel := ParallelGateway{Gateway{FlowNode: FlowNode{BaseHandlerType: delegate.BaseHandlerType(ParallelGateway{}), IncomingFlow: make([]delegate.FlowElement, 0), OutgoingFlow: make([]delegate.FlowElement, 0)}}}
	decoder.DecodeElement(&parallel, &token)
	return &parallel
}
