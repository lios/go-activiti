package converter

import (
	. "encoding/xml"
	"github.com/lios/go-activiti/engine/contanst"
	. "github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type UserTaskXMLConverter struct {
	BpmnXMLConverter
}

func (user UserTaskXMLConverter) GetXMLElementName() string {
	return contanst.ELEMENT_TASK_USER
}

func (user UserTaskXMLConverter) ConvertXMLToElement(decoder *Decoder, token StartElement, model *BpmnModel, activeProcess *Process) BaseElement {

	userTask := UserTask{FlowNode: FlowNode{BaseHandlerType: BaseHandlerType(UserTask{}), IncomingFlow: make([]FlowElement, 0), OutgoingFlow: make([]FlowElement, 0)}}
	decoder.DecodeElement(&userTask, &token)
	return &userTask
}
