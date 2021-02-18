package converter

import (
	"bytes"
	. "encoding/xml"
	. "github.com/lios/go-activiti/engine/contanst"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/converter/parser"
	"github.com/lios/go-activiti/engine/impl/delegate"
	"github.com/lios/go-activiti/logger"
)

var convertersToBpmnMap map[string]BaseBpmnXMLConverter

func init() {
	convertersToBpmnMap = make(map[string]BaseBpmnXMLConverter, 0)
	AddConverter(UserTaskXMLConverter{BpmnXMLConverter{BaseBpmnXMLConverter(UserTaskXMLConverter{})}})
	AddConverter(SequenceFlowXMLConverter{BpmnXMLConverter{BaseBpmnXMLConverter(SequenceFlowXMLConverter{})}})
	AddConverter(StartEventXMLConverter{BpmnXMLConverter{BaseBpmnXMLConverter(StartEventXMLConverter{})}})
	AddConverter(EndEventXMLConverter{BpmnXMLConverter{BaseBpmnXMLConverter(EndEventXMLConverter{})}})
	AddConverter(InclusiveGatewayXMLConverter{BpmnXMLConverter{BaseBpmnXMLConverter(InclusiveGatewayXMLConverter{})}})
	AddConverter(ExclusiveGatewayXMLConverter{BpmnXMLConverter{BaseBpmnXMLConverter(ExclusiveGatewayXMLConverter{})}})
	AddConverter(ParallelGatewayXMLConverter{BpmnXMLConverter{BaseBpmnXMLConverter(ParallelGatewayXMLConverter{})}})

}

type BpmnXMLConverter struct {
	BaseBpmnXMLConverter
}

func AddConverter(converter BaseBpmnXMLConverter) {
	convertersToBpmnMap[converter.GetXMLElementName()] = converter
}

func (BpmnXMLConverter) ConvertToBpmnModel(byte []byte) *BpmnModel {
	reader := bytes.NewReader(byte)
	// 创建带缓存的 Reader
	decoder := NewDecoder(reader)
	decoder.Token()
	model := &BpmnModel{make([]*Process, 0)}
	bpmnDecoder(decoder, model)

	for _, process := range model.GetMainProcess() {
		processFlowElements(process.FlowElementList, process)
	}
	return model
}

func bpmnDecoder(decoder *Decoder, model *BpmnModel) {
	processParser := parser.ProcessParser{}
	var process *Process
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case StartElement:
			name := token.Name.Local
			if name == ELEMENT_PROCESS {
				process = processParser.Parse(decoder, token, model)
			}
			converter, ok := convertersToBpmnMap[name]
			if ok {
				converter.convertToBpmnModel(decoder, token, model, process)
			}
		case EndElement:
			logger.Warn("xml parser end")
		}
	}
}
func (bpmnXMLConverter BpmnXMLConverter) convertToBpmnModel(decoder *Decoder, token StartElement, model *BpmnModel, activeProcess *Process) {
	parsedElement := bpmnXMLConverter.BaseBpmnXMLConverter.ConvertXMLToElement(decoder, token, model, activeProcess)
	flowElement := parsedElement.(delegate.FlowElement)
	activeProcess.AddFlowElement(flowElement)
}

func processFlowElements(flowElementList []delegate.FlowElement, parentScope delegate.BaseElement) {
	for _, flowElement := range flowElementList {
		sequenceFlow, ok := flowElement.(*SequenceFlow)
		if !ok {
			continue
		}
		sourceNode := getFlowNodeFromScope(sequenceFlow.SourceRef, parentScope)
		if sourceNode != nil {
			outgoing := append(sourceNode.GetOutgoing(), sequenceFlow)
			sourceNode.SetOutgoing(outgoing)
			sequenceFlow.SetSourceFlowElement(sourceNode)
		}

		targetNode := getFlowNodeFromScope(sequenceFlow.TargetRef, parentScope)
		if targetNode != nil {
			ingoing := append(targetNode.GetIncoming(), sequenceFlow)
			targetNode.SetIncoming(ingoing)
			sequenceFlow.SetTargetFlowElement(targetNode)
		}

	}
}
func getFlowNodeFromScope(elementId string, scope delegate.BaseElement) delegate.FlowElement {
	process, ok := scope.(*Process)
	if !ok {
		return nil
	}
	if elementId != "" {
		element := process.GetFlowElement(elementId)
		//flowNode ,ok = element.(FlowNode)
		//if ok {
		//	return flowNode
		//}
		return element
	}
	return nil
}
