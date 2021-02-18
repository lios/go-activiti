package agenda

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
	. "github.com/lios/go-activiti/engine/impl/invocation/behavior"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type TriggerExecutionOperation struct {
	AbstractOperation
}

func (trigger TriggerExecutionOperation) Run() (err error) {
	element := trigger.getCurrentFlowElement(trigger.Execution)
	flowNode := element.(*model.FlowNode)
	behavior := flowNode.GetBehavior()
	operation := behavior.(TriggerableActivityBehavior)
	operation.Trigger(trigger.Execution)
	return err
}

func (trigger TriggerExecutionOperation) getCurrentFlowElement(execut entity.ExecutionEntity) delegate.FlowElement {
	currentFlowElement := execut.GetCurrentFlowElement()
	if currentFlowElement != nil {
		return currentFlowElement
	}
	return nil
}
