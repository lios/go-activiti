package agenda

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
	"github.com/lios/go-activiti/engine/impl/interceptor"
	"github.com/lios/go-activiti/engine/impl/utils"
)

type TakeOutgoingSequenceFlowsOperation struct {
	AbstractOperation
	EvaluateConditions bool
}

func (task TakeOutgoingSequenceFlowsOperation) Run() (err error) {
	currentFlowElement := task.getCurrentFlowElement()
	_, ok := currentFlowElement.(*model.SequenceFlow)
	if ok {
		task.handleSequenceFlow()
	} else {
		err = task.handleFlowNode()
	}
	return err
}

//处理节点
func (task TakeOutgoingSequenceFlowsOperation) handleFlowNode() (err error) {
	execution := task.Execution
	currentFlowElement := task.Execution.GetCurrentFlowElement()
	err = task.handleActivityEnd(currentFlowElement)
	if err != nil {
		return err
	}
	gateway, ok := currentFlowElement.(*model.Gateway)
	var defaultSequenceFlowId = ""
	if ok {
		defaultSequenceFlowId = gateway.DefaultFlow
	}
	flowElements := currentFlowElement.GetOutgoing()
	var outgoingSequenceFlows = make([]delegate.FlowElement, 0)
	if len(flowElements) > 0 {
		for _, flowElement := range flowElements {
			sequenceFlow := (flowElement).(*model.SequenceFlow)
			if !task.EvaluateConditions || utils.HasTrueCondition(*sequenceFlow, execution) {
				outgoingSequenceFlows = append(outgoingSequenceFlows, flowElement)
			}
		}
		if outgoingSequenceFlows != nil && len(outgoingSequenceFlows) == 0 {
			if defaultSequenceFlowId != "" {
				for _, flowElement := range flowElements {
					if defaultSequenceFlowId == (flowElement).GetId() {
						outgoingSequenceFlows = append(outgoingSequenceFlows, flowElement)
					}
				}
			}
		}
	}

	if len(outgoingSequenceFlows) == 0 {
		if flowElements == nil || len(flowElements) == 0 {
			interceptor.GetAgenda().Agenda.PlanEndExecutionOperation(execution)
		} else {
			panic("No outgoing sequence flow of element")
		}
	} else {
		for _, outgoingExecution := range outgoingSequenceFlows {
			execution.SetCurrentFlowElement(outgoingExecution)
			interceptor.GetAgenda().Agenda.PlanContinueProcessOperation(execution)
		}
	}
	return err
}

//处理连线
func (task TakeOutgoingSequenceFlowsOperation) handleSequenceFlow() {
	interceptor.GetAgenda().Agenda.PlanContinueProcessOperation(task.Execution)
}

//获取当前活动节点
func (task TakeOutgoingSequenceFlowsOperation) getCurrentFlowElement() delegate.FlowElement {
	execution := task.Execution
	currentFlowElement := execution.GetCurrentFlowElement()
	if currentFlowElement != nil {
		return currentFlowElement
	}
	return nil
}

func (task TakeOutgoingSequenceFlowsOperation) handleActivityEnd(element delegate.FlowElement) (err error) {
	//historicActinstManager := manager.DataManager{}.HistoricActinstDataManager
	//err = historicActinstManager.RecordTaskCreated(element, task.Execution)
	return err
}
