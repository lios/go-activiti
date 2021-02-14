package agenda

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/invocation"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type ContinueProcessOperation struct {
	AbstractOperation
}

func (cont *ContinueProcessOperation) Run() (err error) {
	element := cont.Execution.GetCurrentFlowElement()
	if element != nil {
		flow := element.(*model.FlowNode)
		if true {
			err = cont.continueThroughFlowNode(*flow)
		} else {
			sequenceFlow, _ := element.(*model.SequenceFlow)
			cont.continueThroughSequenceFlow(*sequenceFlow)
		}
	}
	return err
}

func (cont *ContinueProcessOperation) continueThroughSequenceFlow(sequenceFlow model.SequenceFlow) {
	flowElement := sequenceFlow.TargetFlowElement
	cont.Execution.SetCurrentFlowElement(flowElement)
	invocation.GetAgenda().PlanContinueProcessOperation(cont.Execution)
}

func (cont *ContinueProcessOperation) continueThroughFlowNode(element model.FlowNode) (err error) {
	dataManager := entity.GetHistoricActivityInstanceEntityManager()
	dataManager.RecordActivityStart(cont.Execution)
	behavior := element.GetBehavior()
	if behavior != nil {
		err = behavior.Execute(cont.Execution)
	} else {
		invocation.GetAgenda().PlanTakeOutgoingSequenceFlowsOperation(cont.Execution, true)
	}
	return err
}
