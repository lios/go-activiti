package parse

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	"github.com/lios/go-activiti/engine/impl/bpmn/parse/handler"
	"github.com/lios/go-activiti/logger"
)

type BpmnParseHandlers struct {
	ParseHandlers map[bpmn.BaseElement][]handler.BpmnParseHandler
}

func NewBpmnParseHandlers() BpmnParseHandlers {
	return BpmnParseHandlers{ParseHandlers: make(map[bpmn.BaseElement][]handler.BpmnParseHandler, 0)}
}

func (bpmnParseHandlers BpmnParseHandlers) AddHandlers(name bpmn.BaseElement, bpmnParseHandler []handler.BpmnParseHandler) {
	bpmnParseHandlers.ParseHandlers[name] = bpmnParseHandler
}
func (bpmnParseHandlers BpmnParseHandlers) ParseElement(bpmnParse *BpmnParse, element bpmn.BaseElement) {
	flowElement, ok := element.(bpmn.FlowElement)
	if ok {
		bpmnParse.SetCurrentFlowElement(flowElement)
	}
	handlers := bpmnParseHandlers.ParseHandlers[element]
	if handlers != nil && len(handlers) > 0 {
		for _, handler := range handlers {
			handler.Parse(bpmnParse, element)
		}
	} else {
		logger.Warn("Could not find matching parse handler for + " + element.GetId() + " this is likely a bug.")
	}

}
