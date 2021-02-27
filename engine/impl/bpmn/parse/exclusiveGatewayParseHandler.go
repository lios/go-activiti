package parse

import (
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type ExclusiveGatewayParseHandler struct {
	AbstractActivityBpmnParseHandler
}

func (exclusiveGatewayParseHandler ExclusiveGatewayParseHandler) GetHandledType() string {
	return ExclusiveGateway{}.GetType()
}

func (exclusiveGatewayParseHandler ExclusiveGatewayParseHandler) ExecuteParse(bpmnParse BpmnParse, baseElement delegate.BaseElement) {
	exclusiveGateway := baseElement.(*ExclusiveGateway)
	exclusiveGateway.SetBehavior(bpmnParse.ActivityBehaviorFactory.CreateExclusiveGatewayActivityBehavior(*exclusiveGateway))
}
