package behavior

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/invocation"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
	"github.com/lios/go-activiti/engine/impl/utils"
)

type ExclusiveGatewayActivityBehavior struct {
}

//排他网关
func (exclusive ExclusiveGatewayActivityBehavior) Execute(execution entity.ExecutionEntity) (err error) {
	err = exclusive.Leave(execution)
	return err
}

func (exclusive ExclusiveGatewayActivityBehavior) Leave(execution entity.ExecutionEntity) (err error) {
	element := execution.GetCurrentFlowElement()
	exclusiveGateway, ok := element.(*model.ExclusiveGateway)
	var outgoingSequenceFlow bpmn.FlowElement
	var defaultSequenceFlow bpmn.FlowElement
	if ok {
		defaultSequenceFlowId := exclusiveGateway.DefaultFlow
		sequenceFlowIterator := exclusiveGateway.GetOutgoing()
		if sequenceFlowIterator != nil && len(sequenceFlowIterator) > 0 {
			for _, sequenceFlow := range sequenceFlowIterator {
				flow := (sequenceFlow).(*model.SequenceFlow)
				conditionEvaluatesToTrue := utils.HasTrueCondition(*flow, execution)
				if conditionEvaluatesToTrue && defaultSequenceFlowId != "" && defaultSequenceFlowId != flow.Id {
					outgoingSequenceFlow = sequenceFlow
				}
				if defaultSequenceFlowId != "" && defaultSequenceFlowId == flow.Id {
					defaultSequenceFlow = sequenceFlow
				}
			}

		}
	}
	if outgoingSequenceFlow != nil {
		execution.SetCurrentFlowElement(outgoingSequenceFlow)
	} else {
		if defaultSequenceFlow != nil {
			execution.SetCurrentFlowElement(defaultSequenceFlow)
		}
	}
	//执行出口逻辑，设置条件判断
	invocation.GetAgenda().PlanTakeOutgoingSequenceFlowsOperation(execution, true)
	return nil
}
