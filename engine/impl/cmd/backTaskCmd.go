package cmd

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/cfg"
	. "github.com/lios/go-activiti/engine/impl/persistence/entity"
	"github.com/lios/go-activiti/engine/impl/utils"
	"github.com/lios/go-activiti/engine/interceptor"
	"github.com/pborman/uuid"
)

type BackTaskCmd struct {
	NeedsActiveTaskCmd
	Comment      string
	TargetFlowId string
}

func (backTaskCmd BackTaskCmd) TaskExecute(command interceptor.CommandContext, entity TaskEntity) (interface{}, error) {

	task := entity.(TaskEntityImpl)
	execution := GetExecutionEntityManager().FindById(task.ProcessInstanceId)
	processUtils := utils.ProcessDefinitionUtil{}
	process := processUtils.GetProcess(execution.GetProcessDefineId())
	currentTask := process.GetFlowElement(execution.GetCurrentActivityId())
	sourceElement := currentTask.GetOutgoing()
	targetFlowElement := process.GetFlowElement(backTaskCmd.TargetFlowId)
	sequenceFlows := createTask(targetFlowElement, currentTask.GetId(), backTaskCmd.TargetFlowId)
	currentTask.SetOutgoing(sequenceFlows)
	_, err := cfg.GetCommandExecutorImpl().Exe(CompleteCmd{NeedsActiveTaskCmd: NeedsActiveTaskCmd{TaskId: backTaskCmd.TaskId}, Variables: nil, LocalScope: true})
	if err != nil {
		return false, nil
	}
	currentTask.SetOutgoing(sourceElement)
	return true, err
}

func createTask(element bpmn.FlowElement, sourceRef, targetRef string) []bpmn.FlowElement {
	sequenceFlow := model.SequenceFlow{}
	flow := model.FlowNode{}
	sequenceFlow.FlowNode = &flow
	sequenceFlow.Id = uuid.New()
	sequenceFlow.SourceRef = sourceRef
	sequenceFlow.TargetRef = targetRef
	sequenceFlow.SetTargetFlowElement(element)
	flowElement := make([]bpmn.FlowElement, 0)
	flowElement = append(flowElement, sequenceFlow)
	return flowElement
}
