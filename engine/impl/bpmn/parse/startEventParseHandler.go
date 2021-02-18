package parse

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type StartEventParseHandler struct {
	AbstractActivityBpmnParseHandler
}

func (startEventParseHandler StartEventParseHandler) GetHandledType() string {
	return model.StartEvent{}.GetType()
}

func (startEventParseHandler StartEventParseHandler) ExecuteParse(bpmnParse BpmnParse, flow delegate.BaseElement) {

}
