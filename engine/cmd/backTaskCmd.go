package cmd

import (
	. "github.com/lios/go-activiti/engine"
	"github.com/lios/go-activiti/engine/behavior"
	. "github.com/lios/go-activiti/engine/manager"
	"github.com/pborman/uuid"
)

type BackTaskCmd struct {
	TaskId       int
	TargetFlowId string
	Comment      string
}

func (backTaskCmd BackTaskCmd) Execute(interceptor behavior.CommandContext) (interface{}, error) {
	manager := GetTaskManager()
	task, err := manager.FindById(backTaskCmd.TaskId)
	if err != nil {
		return task, err
	}
	processUtils := behavior.ProcessUtils{}
	currentTask, err := processUtils.GetCurrentTask(backTaskCmd.TaskId)
	if err != nil {
		return false, nil
	}
	sourceElement := currentTask.GetOutgoing()
	targetFlowElement, err := processUtils.GetFlowElement(backTaskCmd.TargetFlowId)
	sequenceFlows := createTask(&targetFlowElement, sourceElement, currentTask.GetId(), backTaskCmd.TargetFlowId)
	currentTask.SetOutgoing(sequenceFlows)
	_, err = behavior.GetServiceImpl().CommandExecutor.Exe(CompleteCmd{backTaskCmd.TaskId, nil, true})
	if err != nil {
		return false, nil
	}
	currentTask.SetOutgoing(sourceElement)
	return true, err
}

func createTask(element *FlowElement, sourceElement []*FlowElement, sourceRef, targetRef string) []*FlowElement {
	elements := make([]*FlowElement, 0)

	elements = append(elements, element)
	sequenceFlow := SequenceFlow{}
	flow := Flow{}
	sequenceFlow.Flow = &flow
	sequenceFlow.Id = uuid.New()
	sequenceFlow.SourceRef = sourceRef
	sequenceFlow.TargetRef = targetRef
	sequenceFlow.SetOutgoing(elements)
	flowElement := make([]*FlowElement, 0)
	//flowElement = append(flowElement,sequenceFlow)
	return flowElement
}
