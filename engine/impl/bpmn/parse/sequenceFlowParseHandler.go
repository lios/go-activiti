package parse

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type SequenceFlowParseHandler struct {
	AbstractActivityBpmnParseHandler
}

func (sequenceFlowParseHandler SequenceFlowParseHandler) GetHandledType() string {
	return model.SequenceFlow{}.GetType()
}

func (sequenceFlowParseHandler SequenceFlowParseHandler) ExecuteParse(bpmnParse BpmnParse, flow delegate.BaseElement) {
	process := bpmnParse.CurrentProcess
	flowElement := flow.(*model.SequenceFlow)
	flowElement.SetSourceFlowElement(process.GetFlowElement(flowElement.SourceRef))
	flowElement.SetTargetFlowElement(process.GetFlowElement(flowElement.TargetRef))
}
