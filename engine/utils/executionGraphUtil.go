package utils

import . "github.com/lios/go-activiti/engine"

type ExecutionGraphUtil struct {
}

func IsReachable(process Process, sourceElementId string, targetElementId string) bool {
	sourceFlowElement := process.FlowMap[sourceElementId]
	sourceFlow, ok := sourceFlowElement.(SequenceFlow)
	if !ok {
		element := sourceFlow.GetTargetFlowElement()
		flow, _ := (*element).(SequenceFlow)
		sourceFlow = flow
	}

	targetFlowElement := process.FlowMap[targetElementId]
	targetFlow, ok := targetFlowElement.(SequenceFlow)
	if !ok {
		element := targetFlow.GetTargetFlowElement()
		flow, _ := (*element).(SequenceFlow)
		targetFlow = flow
	}
	var visitedElements = make(map[string]FlowElement, 0)
	return isReachable(process, sourceFlow, targetFlow, visitedElements)

}

func isReachable(process Process, sourceElement FlowElement, targetElement FlowElement, visitedElements map[string]FlowElement) bool {
	if sourceElement.GetId() == targetElement.GetId() {
		return true
	}
	visitedElements[sourceElement.GetId()] = sourceElement
	outgoing := sourceElement.GetOutgoing()
	if outgoing != nil && len(outgoing) > 0 {
		for _, value := range outgoing {
			sequenceFlowTarget := (*value).GetTargetFlowElement()
			if sequenceFlowTarget != nil && visitedElements[(*sequenceFlowTarget).GetId()] != nil {
				var reachable = isReachable(process, *sequenceFlowTarget, targetElement, visitedElements)
				if reachable {
					return true
				}
			}
		}
	}
	return false
}
