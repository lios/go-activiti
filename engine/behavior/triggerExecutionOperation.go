package behavior

import (
	"github.com/lios/go-activiti/engine"
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

func (trigger TriggerExecutionOperation) getCurrentFlowElement(execut engine.ExecutionEntity) engine.FlowElement {
	currentFlowElement := execut.GetCurrentFlowElement()
	if currentFlowElement != nil {
		return currentFlowElement
	}
	return nil
}
