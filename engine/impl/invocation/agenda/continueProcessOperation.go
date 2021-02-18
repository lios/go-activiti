package agenda

import (
	. "github.com/lios/go-activiti/engine/impl/delegate"
	"github.com/lios/go-activiti/engine/impl/interceptor"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type ContinueProcessOperation struct {
	AbstractOperation
}

func (cont *ContinueProcessOperation) Run() (err error) {
	element := cont.Execution.GetCurrentFlowElement()
	if element != nil {
		if element.GetOutgoing() != nil {
			err = cont.continueThroughFlowNode(element)
		} else {
			cont.continueThroughSequenceFlow(element)
		}
	}
	return err
}

func (cont *ContinueProcessOperation) continueThroughSequenceFlow(sequenceFlow FlowElement) {
	flowElement := sequenceFlow.GetTargetFlowElement()
	cont.Execution.SetCurrentFlowElement(flowElement)
	interceptor.GetAgenda().Agenda.PlanContinueProcessOperation(cont.Execution)
}

func (cont *ContinueProcessOperation) continueThroughFlowNode(element FlowElement) (err error) {
	dataManager := entity.GetHistoricActivityInstanceEntityManager()
	dataManager.RecordActivityStart(cont.Execution)
	behavior := element.GetBehavior()
	if behavior != nil {
		err = behavior.Execute(cont.Execution)
	} else {
		interceptor.GetAgenda().Agenda.PlanTakeOutgoingSequenceFlowsOperation(cont.Execution, true)
	}
	return err
}
