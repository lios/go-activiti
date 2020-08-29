package behavior

import (
	"github.com/lios/go-activiti/engine"
)

type ContinueProcessOperation struct {
	AbstractOperation
}

func (cont *ContinueProcessOperation) Run() (err error) {
	element := cont.Execution.GetCurrentFlowElement()
	if element != nil {
		flow, ok := element.(engine.SequenceFlow)
		if !ok {
			err = cont.continueThroughFlowNode(element)
		} else {
			cont.continueThroughSequenceFlow(flow)
		}
	}
	return err
}

func (cont *ContinueProcessOperation) continueThroughSequenceFlow(sequenceFlow engine.SequenceFlow) {
	flowElement := sequenceFlow.TargetFlowElement
	cont.Execution.SetCurrentFlowElement(*flowElement)
	GetAgenda().PlanContinueProcessOperation(cont.Execution)
}

func (cont *ContinueProcessOperation) continueThroughFlowNode(element engine.FlowElement) (err error) {
	historicActinstManager := GetHistoricActinstManager()
	historicActinstManager.RecordActivityStart(cont.Execution)
	behavior := element.GetBehavior()
	if behavior != nil {
		err = behavior.Execute(cont.Execution)
	} else {
		GetAgenda().PlanTakeOutgoingSequenceFlowsOperation(cont.Execution, true)
	}
	return err
}
