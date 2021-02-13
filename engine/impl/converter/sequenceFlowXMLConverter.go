package converter

import (
	. "encoding/xml"
	"github.com/lios/go-activiti/engine/contanst"
	. "github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type SequenceFlowXMLConverter struct {
	BpmnXMLConverter
}

func (sequence SequenceFlowXMLConverter) GetXMLElementName() string {
	return contanst.ELEMENT_SEQUENCE_FLOW
}

func (sequence SequenceFlowXMLConverter) ConvertXMLToElement(decoder *Decoder, token StartElement, model *BpmnModel, activeProcess *Process) BaseElement {
	sequenceFlow := SequenceFlow{FlowNode: FlowNode{BaseHandlerType: BaseHandlerType(EndEvent{})}}
	decoder.DecodeElement(&sequenceFlow, &token)
	return &sequenceFlow
}
