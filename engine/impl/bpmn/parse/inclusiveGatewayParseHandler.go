package parse

import (
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type InclusiveGatewayParseHandler struct {
	AbstractActivityBpmnParseHandler
}

func (inclusiveGatewayParseHandler InclusiveGatewayParseHandler) GetHandledType() string {
	return InclusiveGateway{}.GetType()
}

func (inclusiveGatewayParseHandler InclusiveGatewayParseHandler) ExecuteParse(bpmnParse BpmnParse, baseElement delegate.BaseElement) {
	inclusiveGateway := baseElement.(*InclusiveGateway)
	inclusiveGateway.SetBehavior(bpmnParse.ActivityBehaviorFactory.CreateInclusiveGatewayActivityBehavior(*inclusiveGateway))
}
