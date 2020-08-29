package behavior

import (
	"github.com/lios/go-activiti/engine"
	"github.com/lios/go-activiti/engine/utils"
)

type ExclusiveGatewayActivityBehavior struct {
}

//排他网关
func (exclusive ExclusiveGatewayActivityBehavior) Execute(execution engine.ExecutionEntity) (err error) {
	err = exclusive.Leave(execution)
	return err
}

func (exclusive ExclusiveGatewayActivityBehavior) Leave(execution engine.ExecutionEntity) (err error) {
	element := execution.GetCurrentFlowElement()
	exclusiveGateway, ok := element.(engine.ExclusiveGateway)
	var outgoingSequenceFlow *engine.FlowElement
	var defaultSequenceFlow *engine.FlowElement
	if ok {
		defaultSequenceFlowId := exclusiveGateway.DefaultFlow
		sequenceFlowIterator := exclusiveGateway.GetOutgoing()
		if sequenceFlowIterator != nil && len(sequenceFlowIterator) > 0 {
			for _, sequenceFlow := range sequenceFlowIterator {
				flow := (*sequenceFlow).(engine.SequenceFlow)
				conditionEvaluatesToTrue := utils.HasTrueCondition(flow, execution)
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
		execution.SetCurrentFlowElement(*outgoingSequenceFlow)
	} else {
		if defaultSequenceFlow != nil {
			execution.SetCurrentFlowElement(*defaultSequenceFlow)
		}
	}
	//执行出口逻辑，设置条件判断
	GetAgenda().PlanTakeOutgoingSequenceFlowsOperation(execution, true)
	return nil
}
