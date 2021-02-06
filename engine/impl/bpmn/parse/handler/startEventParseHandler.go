package handler

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/bpmn/parse"
)

type StartEventParseHandler struct {
	AbstractActivityBpmnParseHandler
}

func (startEventParseHandler StartEventParseHandler) GetHandledType() bpmn.BaseElement {
	return model.StartEvent{}
}

func (startEventParseHandler StartEventParseHandler) ExecuteParse(bpmnParse parse.BpmnParse, flow bpmn.BaseElement) {

}
