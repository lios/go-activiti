package agenda

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/manager"
)

type ContinueProcessOperation struct {
	AbstractOperation
}

func (cont *ContinueProcessOperation) Run() (err error) {
	element := cont.Execution.GetCurrentFlowElement()
	if element != nil {
		flow, ok := element.(model.SequenceFlow)
		if !ok {
			err = cont.continueThroughFlowNode(element)
		} else {
			cont.continueThroughSequenceFlow(flow)
		}
	}
	return err
}

func (cont *ContinueProcessOperation) continueThroughSequenceFlow(sequenceFlow model.SequenceFlow) {
	flowElement := sequenceFlow.TargetFlowElement
	cont.Execution.SetCurrentFlowElement(flowElement)
	PlanContinueProcessOperation(cont.Execution)
}

func (cont *ContinueProcessOperation) continueThroughFlowNode(element model.FlowElement) (err error) {
	historicActinstManager := manager.GetDataManager().HistoricActinstDataManager
	historicActinstManager.RecordActivityStart(cont.Execution)
	behavior := element.GetBehavior()
	if behavior != nil {
		err = behavior.Execute(cont.Execution)
	} else {
		PlanTakeOutgoingSequenceFlowsOperation(cont.Execution, true)
	}
	return err
}