package agenda

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	. "github.com/lios/go-activiti/engine/impl/invocation/behavior"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type TriggerExecutionOperation struct {
	AbstractOperation
}

func (trigger TriggerExecutionOperation) Run() (err error) {
	element := trigger.getCurrentFlowElement(trigger.Execution)
	behavior := element.GetBehavior()
	operation := behavior.(TriggerableActivityBehavior)
	operation.Trigger(trigger.Execution)
	return err
}

func (trigger TriggerExecutionOperation) getCurrentFlowElement(execut entity.ExecutionEntity) model.FlowElement {
	currentFlowElement := execut.GetCurrentFlowElement()
	if currentFlowElement != nil {
		return currentFlowElement
	}
	return nil
}
