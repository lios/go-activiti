package parser

import (
	. "encoding/xml"
	. "github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type ProcessParser struct {
}

func (ProcessParser ProcessParser) Parse(decoder *Decoder, token StartElement, model *BpmnModel) *Process {
	process := Process{FlowNode: FlowNode{BaseHandlerType: BaseHandlerType(Process{})}, FlowElementList: make([]FlowElement, 0), FlowElementMap: make(map[string]FlowElement, 0)}
	attrs := token.Attr
	tem := make(map[string]string, 0)
	for _, attr := range attrs {
		tem[attr.Name.Local] = attr.Value
	}
	//decoder.DecodeElement(process, &token)
	process.Id = tem["id"]
	model.AddProcess(&process)
	return &process
}
