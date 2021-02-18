package parser

import (
	. "encoding/xml"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type ProcessParser struct {
}

func (ProcessParser ProcessParser) Parse(decoder *Decoder, token StartElement, model *BpmnModel) *Process {
	process := Process{FlowNode: FlowNode{BaseHandlerType: delegate.BaseHandlerType(Process{})}, FlowElementList: make([]delegate.FlowElement, 0), FlowElementMap: make(map[string]delegate.FlowElement, 0)}
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
