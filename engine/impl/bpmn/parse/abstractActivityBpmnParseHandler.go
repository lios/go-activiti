package parse

import (
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type AbstractActivityBpmnParseHandler struct {
	AbstractBpmnParseHandler
}

func (abstractBpmnParse AbstractActivityBpmnParseHandler) Parse(bpmnParse *BpmnParse, element delegate.BaseElement) {
	abstractBpmnParse.AbstractBpmnParseHandler.Parse(bpmnParse, element)
	activity, ok := element.(Activity)
	if ok && activity.LoopCharacteristics != nil {
		abstractBpmnParse.createMultiInstanceLoopCharacteristics(*bpmnParse, activity)
	}
}
func (abstractBpmnParse AbstractActivityBpmnParseHandler) createMultiInstanceLoopCharacteristics(bpmnParse BpmnParse, element delegate.BaseElement) {

}
