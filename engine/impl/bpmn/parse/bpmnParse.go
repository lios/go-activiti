package parse

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/bpmn/parse/factory"
)

type BpmnParse struct {
	Name                    string
	ActivityBehaviorFactory factory.ActivityBehaviorFactory
	BpmnParserHandlers      BpmnParseHandlers
	bpmnModel               model.BpmnModel
}

func (bpmnParse *BpmnParse) Execute() {
	bpmnParse.applyParseHandlers()
}

func (bpmnParse *BpmnParse) applyParseHandlers() {
	for _, process := range bpmnParse.bpmnModel.GetMainProcess() {
		bpmnParse.BpmnParserHandlers.ParseElement(bpmnParse, process)
	}

}
