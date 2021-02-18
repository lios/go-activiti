package parse

import (
	"github.com/lios/go-activiti/engine/impl/delegate"
	"github.com/lios/go-activiti/logger"
)

type BpmnParseHandlers struct {
	ParseHandlers map[string][]BpmnParseHandler
}

func NewBpmnParseHandlers() BpmnParseHandlers {
	return BpmnParseHandlers{ParseHandlers: make(map[string][]BpmnParseHandler, 0)}
}
func (bpmnParseHandlers BpmnParseHandlers) AddHandlers(handlers []BpmnParseHandler) {
	for _, handler := range handlers {
		bpmnParseHandlers.AddHandler(handler)
	}
}

func (bpmnParseHandlers BpmnParseHandlers) AddHandler(bpmnParseHandler BpmnParseHandler) {
	handledTypes := bpmnParseHandler.GetHandledTypes()
	for _, handledType := range handledTypes {
		_, ok := bpmnParseHandlers.ParseHandlers[handledType]
		if !ok {
			parseHandlers := make([]BpmnParseHandler, 0)
			bpmnParseHandlers.ParseHandlers[handledType] = parseHandlers
		}
		bpmnParseHandlers.ParseHandlers[handledType] = append(bpmnParseHandlers.ParseHandlers[handledType], bpmnParseHandler)

	}
}
func (bpmnParseHandlers BpmnParseHandlers) ParseElement(bpmnParse *BpmnParse, element delegate.BaseElement) {
	flowElement, ok := element.(delegate.FlowElement)
	if ok {
		bpmnParse.SetCurrentFlowElement(flowElement)
	}
	handlerType := flowElement.GetHandlerType()
	handlers := bpmnParseHandlers.ParseHandlers[handlerType]
	if handlers != nil && len(handlers) > 0 {
		for _, handler := range handlers {
			handler.Parse(bpmnParse, element)
		}
	} else {
		logger.Warn("Could not find matching parse handler for + " + element.GetId() + " this is likely a bug.")
	}

}
