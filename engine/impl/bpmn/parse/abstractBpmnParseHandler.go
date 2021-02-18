package parse

import (
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type AbstractBpmnParseHandler struct {
	ParseHandler
}

func (abstractBpmnParse AbstractBpmnParseHandler) GetHandledTypes() []string {
	types := make([]string, 0)
	types = append(types, abstractBpmnParse.ParseHandler.GetHandledType())
	return types
}
func (abstractBpmnParse AbstractBpmnParseHandler) Parse(bpmnParse *BpmnParse, element delegate.BaseElement) {
	abstractBpmnParse.ExecuteParse(*bpmnParse, element)
}
