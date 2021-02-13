package converter

import (
	. "encoding/xml"
	"github.com/lios/go-activiti/engine/contanst"
	. "github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type EndEventXMLConverter struct {
	BpmnXMLConverter
}

func (endEvent EndEventXMLConverter) GetXMLElementName() string {
	return contanst.ELEMENT_EVENT_END
}
func (endEvent EndEventXMLConverter) ConvertXMLToElement(decoder *Decoder, token StartElement, model *BpmnModel, activeProcess *Process) BaseElement {
	end := EndEvent{FlowNode: FlowNode{BaseHandlerType: BaseHandlerType(EndEvent{}), IncomingFlow: make([]FlowElement, 0), OutgoingFlow: make([]FlowElement, 0)}}
	decoder.DecodeElement(&end, &token)
	return &end
}
