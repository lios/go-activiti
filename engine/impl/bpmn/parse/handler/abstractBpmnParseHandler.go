package handler

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	. "github.com/lios/go-activiti/engine/impl/bpmn/parse"
)

type AbstractBpmnParseHandler struct {
	ParseHandler
}

func (abstractBpmnParse AbstractBpmnParseHandler) GetHandledTypes() []bpmn.BaseElement {
	types := make([]bpmn.BaseElement, 0)
	types = append(types, abstractBpmnParse.GetHandledType())
	return types
}
func (abstractBpmnParse AbstractBpmnParseHandler) Parse(bpmnParse *BpmnParse, element bpmn.BaseElement) {
	abstractBpmnParse.ExecuteParse(*bpmnParse, element)
}
