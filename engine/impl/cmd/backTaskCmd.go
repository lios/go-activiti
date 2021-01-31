package cmd

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/utils"
	"github.com/lios/go-activiti/engine/interceptor"
	"github.com/pborman/uuid"
)

type BackTaskCmd struct {
	TaskId       int
	TargetFlowId string
	Comment      string
}

func (backTaskCmd BackTaskCmd) Execute(commandContext interceptor.CommandContext) (interface{}, error) {
	manager := commandContext.ProcessEngineConfiguration.TaskDataManager
	task, err := manager.FindById(backTaskCmd.TaskId)
	if err != nil {
		return task, err
	}
	processUtils := utils.ProcessUtils{}
	currentTask, err := processUtils.GetCurrentTask(backTaskCmd.TaskId)
	if err != nil {
		return false, nil
	}
	sourceElement := currentTask.GetOutgoing()
	targetFlowElement, err := processUtils.GetFlowElement(backTaskCmd.TargetFlowId)
	sequenceFlows := createTask(targetFlowElement, currentTask.GetId(), backTaskCmd.TargetFlowId)
	currentTask.SetOutgoing(sequenceFlows)
	_, err = commandContext.ProcessEngineConfiguration.CommandExecutor.Exe(CompleteCmd{backTaskCmd.TaskId, nil, true})
	if err != nil {
		return false, nil
	}
	currentTask.SetOutgoing(sourceElement)
	return true, err
}

func createTask(element model.FlowElement, sourceRef, targetRef string) []model.FlowElement {
	sequenceFlow := model.SequenceFlow{}
	flow := model.Flow{}
	sequenceFlow.Flow = &flow
	sequenceFlow.Id = uuid.New()
	sequenceFlow.SourceRef = sourceRef
	sequenceFlow.TargetRef = targetRef
	sequenceFlow.SetTargetFlowElement(element)
	flowElement := make([]model.FlowElement, 0)
	flowElement = append(flowElement, sequenceFlow)
	return flowElement
}
