package cmd

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
	"github.com/lios/go-activiti/engine/impl/interceptor"
	. "github.com/lios/go-activiti/engine/impl/persistence/entity"
	"github.com/lios/go-activiti/engine/impl/utils"
	"github.com/pborman/uuid"
)

type BackTaskCmd struct {
	NeedsActiveTaskCmd
	Comment      string
	TargetFlowId string
}

func (backTaskCmd BackTaskCmd) TaskExecute(command interceptor.CommandContext, entity TaskEntity) (interface{}, error) {

	task := entity.(*TaskEntityImpl)
	execution := GetExecutionEntityManager().FindById(task.ProcessInstanceId)
	processUtils := utils.ProcessDefinitionUtil{}
	process := processUtils.GetProcess(execution.GetProcessDefineId())
	currentTask := process.GetFlowElement(task.GetTaskDefineKey())
	sourceElement := currentTask.GetOutgoing()
	targetFlowElement := process.GetFlowElement(backTaskCmd.TargetFlowId)
	sequenceFlows := createTask(targetFlowElement, currentTask.GetId(), backTaskCmd.TargetFlowId)
	currentTask.SetOutgoing(sequenceFlows)
	_, err := interceptor.GetCommandExecutorImpl().Exe(CompleteCmd{NeedsActiveTaskCmd: NeedsActiveTaskCmd{AbstractTaskCmd: AbstractTaskCmd(CompleteCmd{Variables: nil, LocalScope: true}), TaskId: backTaskCmd.TaskId}})
	if err != nil {
		return false, nil
	}
	currentTask.SetOutgoing(sourceElement)
	return true, err
}

func createTask(element delegate.FlowElement, sourceRef, targetRef string) []delegate.FlowElement {
	sequenceFlow := model.SequenceFlow{}
	flow := model.FlowNode{}
	sequenceFlow.FlowNode = flow
	sequenceFlow.Id = uuid.New()
	sequenceFlow.SourceRef = sourceRef
	sequenceFlow.TargetRef = targetRef
	sequenceFlow.SetTargetFlowElement(element)
	flowElement := make([]delegate.FlowElement, 0)
	flowElement = append(flowElement, &sequenceFlow)
	return flowElement
}
