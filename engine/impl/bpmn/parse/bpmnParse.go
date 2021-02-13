package parse

import (
	. "github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/bpmn/parse/factory"
	"github.com/lios/go-activiti/engine/impl/converter"
	. "github.com/lios/go-activiti/engine/impl/persistence/entity"
	"github.com/lios/go-activiti/logger"
)

type BpmnParse struct {
	Name                    string
	Byte                    []byte
	DeploymentEntity        DeploymentEntity
	ActivityBehaviorFactory factory.ActivityBehaviorFactory
	BpmnParserHandlers      BpmnParseHandlers
	BpmnModel               *BpmnModel
	CurrentFlowElement      FlowElement
	CurrentProcess          *Process
	ProcessDefinitions      []ProcessDefinitionEntity
}

func (bpmnParse BpmnParse) SourceInputStream(byte []byte) BpmnParse {
	bpmnParse.setStreamSource(byte)
	return bpmnParse
}

func (bpmnParse BpmnParse) Deployment(deployment DeploymentEntity) BpmnParse {
	bpmnParse.DeploymentEntity = deployment
	return bpmnParse
}

func (bpmnParse BpmnParse) SourceName(name string) BpmnParse {
	bpmnParse.Name = name
	return bpmnParse
}
func (bpmnParse *BpmnParse) setStreamSource(byte []byte) {
	if byte == nil {
		logger.Error("invalid: multiple sources ")
		panic("invalid: multiple sources ")
	}
	bpmnParse.Byte = byte
}
func (bpmnParse *BpmnParse) Execute() {
	xmlConverter := converter.BpmnXMLConverter{}
	bpmnModel := xmlConverter.ConvertToBpmnModel(bpmnParse.Byte)
	bpmnParse.BpmnModel = bpmnModel
	bpmnParse.applyParseHandlers()
}

func (bpmnParse *BpmnParse) applyParseHandlers() {
	for _, process := range bpmnParse.BpmnModel.GetMainProcess() {
		bpmnParse.CurrentProcess = process
		bpmnParse.BpmnParserHandlers.ParseElement(bpmnParse, process)
	}
}

func (bpmnParse *BpmnParse) SetCurrentFlowElement(currentFlowElement FlowElement) {
	bpmnParse.CurrentFlowElement = currentFlowElement
}

func (bpmnParse *BpmnParse) ProcessFlowElements(flowElements []FlowElement) {
	for _, element := range flowElements {
		bpmnParse.BpmnParserHandlers.ParseElement(bpmnParse, element)
	}
}
