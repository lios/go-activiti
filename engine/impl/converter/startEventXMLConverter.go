package converter

import (
	. "encoding/xml"
	"github.com/lios/go-activiti/engine/contanst"
	"github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type StartEventXMLConverter struct {
	BpmnXMLConverter
}

func (start StartEventXMLConverter) GetXMLElementName() string {
	return contanst.ELEMENT_EVENT_START
}
func (start StartEventXMLConverter) ConvertXMLToElement(decoder *Decoder, token StartElement, model BpmnModel, activeProcess Process) bpmn.BaseElement {
	startEvent := StartEvent{}

	return startEvent
}
