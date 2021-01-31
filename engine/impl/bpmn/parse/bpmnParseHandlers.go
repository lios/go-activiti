package parse

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type BpmnParseHandlers struct {
	ParseHandlers map[model.BaseElement][]BpmnParseHandler
}

func NewBpmnParseHandlers() BpmnParseHandlers {
	return BpmnParseHandlers{ParseHandlers: make(map[model.BaseElement][]BpmnParseHandler, 0)}
}

func (bpmnParse BpmnParseHandlers) AddHandlers(name model.BaseElement, bpmnParseHandler []BpmnParseHandler) {
	bpmnParse.ParseHandlers[name] = bpmnParseHandler
}
func (bpmnParseHandlers BpmnParseHandlers) ParseElement(bpmnParse *BpmnParse, element model.BaseElement) {
	handlers := bpmnParseHandlers.ParseHandlers[element]
	for _, handler := range handlers {
		Parse(bpmnParse, element)
	}
}
