package converter

import (
	. "encoding/xml"
	"github.com/lios/go-activiti/engine/contanst"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type UserTaskXMLConverter struct {
	BpmnXMLConverter
}

func (user UserTaskXMLConverter) GetXMLElementName() string {
	return contanst.ELEMENT_TASK_USER
}

func (user UserTaskXMLConverter) ConvertXMLToElement(decoder *Decoder, token StartElement, model *BpmnModel, activeProcess *Process) delegate.BaseElement {

	userTask := UserTask{FlowNode: FlowNode{BaseHandlerType: delegate.BaseHandlerType(UserTask{}), IncomingFlow: make([]delegate.FlowElement, 0), OutgoingFlow: make([]delegate.FlowElement, 0)}}
	decoder.DecodeElement(&userTask, &token)
	return &userTask
}
