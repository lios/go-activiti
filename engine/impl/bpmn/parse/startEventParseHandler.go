package parse

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type StartEventParseHandler struct {
	AbstractActivityBpmnParseHandler
}

func (startEventParseHandler StartEventParseHandler) GetHandledType() string {
	return model.StartEvent{}.GetType()
}

func (startEventParseHandler StartEventParseHandler) ExecuteParse(bpmnParse BpmnParse, flow bpmn.BaseElement) {

}
