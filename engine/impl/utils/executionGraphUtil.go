package utils

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type ExecutionGraphUtil struct {
}

func IsReachable(processDefinitionId int64, sourceElementId string, targetElementId string) bool {
	processDefinitionUtil := ProcessDefinitionUtil{}
	process := processDefinitionUtil.GetProcess(processDefinitionId)
	sourceFlowElement := process.GetFlowElement(sourceElementId)
	sourceFlow, ok := sourceFlowElement.(*model.SequenceFlow)
	if !ok {
		element := sourceFlow.GetTargetFlowElement()
		flow, _ := (element).(*model.SequenceFlow)
		sourceFlow = flow
	}

	targetFlowElement := process.GetFlowElement(targetElementId)
	targetFlow, ok := targetFlowElement.(*model.SequenceFlow)
	if !ok {
		element := targetFlow.GetTargetFlowElement()
		flow, _ := (element).(*model.SequenceFlow)
		targetFlow = flow
	}
	var visitedElements = make(map[string]bpmn.FlowElement, 0)
	return isReachable(process, sourceFlow, targetFlow, visitedElements)

}

func isReachable(process model.Process, sourceElement bpmn.FlowElement, targetElement bpmn.FlowElement, visitedElements map[string]bpmn.FlowElement) bool {
	if sourceElement.GetId() == targetElement.GetId() {
		return true
	}
	visitedElements[sourceElement.GetId()] = sourceElement
	outgoing := sourceElement.GetOutgoing()
	if outgoing != nil && len(outgoing) > 0 {
		for _, value := range outgoing {
			sequenceFlowTarget := (value).GetTargetFlowElement()
			if sequenceFlowTarget != nil && visitedElements[(sequenceFlowTarget).GetId()] != nil {
				var reachable = isReachable(process, sequenceFlowTarget, targetElement, visitedElements)
				if reachable {
					return true
				}
			}
		}
	}
	return false
}
