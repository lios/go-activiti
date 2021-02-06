package converter

import (
	"bytes"
	. "encoding/xml"
	. "github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/converter/parser"
	"io"
)

var convertersToBpmnMap map[string]BaseBpmnXMLConverter

func init() {
	convertersToBpmnMap = make(map[string]BaseBpmnXMLConverter, 0)
}

type BpmnXMLConverter struct {
	BaseBpmnXMLConverter
}

func (BpmnXMLConverter) AddConverter(converter BaseBpmnXMLConverter) {
	convertersToBpmnMap[converter.GetXMLElementName()] = converter
}
func (BpmnXMLConverter) ConvertToBpmnModel(byte []byte) *BpmnModel {
	model := BpmnModel{}
	processParser := parser.ProcessParser{}
	reader := bytes.NewReader(byte)
	// 创建带缓存的 Reader
	decoder := NewDecoder(reader)
	bpmnModel := new(BpmnModel)
	for t, err := decoder.Token(); err == nil || err == io.EOF; t, err = decoder.Token() {
		switch token := t.(type) {
		case StartElement:
			process := processParser.Parse(decoder, token, *bpmnModel)
			name := token.Name.Local
			converter, ok := convertersToBpmnMap[name]
			if ok {
				converter.convertToBpmnModel(decoder, token, model, process)
			}
		}
	}
	for _, process := range model.GetMainProcess() {
		processFlowElements(process.FlowElementList, process)
	}
	return bpmnModel
}
func (bpmnXMLConverter BpmnXMLConverter) convertToBpmnModel(decoder *Decoder, token StartElement, model BpmnModel, activeProcess Process) {
	parsedElement := bpmnXMLConverter.ConvertXMLToElement(decoder, token, model, activeProcess)
	flowElement := parsedElement.(FlowElement)
	activeProcess.AddFlowElement(flowElement)
}

func processFlowElements(flowElementList []FlowElement, parentScope BaseElement) {
	for _, flowElement := range flowElementList {
		sequenceFlow := flowElement.(SequenceFlow)
		sourceNode := getFlowNodeFromScope(sequenceFlow.SourceRef, parentScope)
		if sourceNode != nil {
			outgoing := append(sourceNode.GetOutgoing(), sequenceFlow)
			sourceNode.SetOutgoing(outgoing)
			sequenceFlow.SetSourceFlowElement(sourceNode)
		}

		targetNode := getFlowNodeFromScope(sequenceFlow.TargetRef, parentScope)
		if targetNode != nil {
			ingoing := append(sourceNode.GetIncoming(), sequenceFlow)
			sourceNode.SetOutgoing(ingoing)
			sequenceFlow.SetTargetFlowElement(targetNode)
		}
	}
}
func getFlowNodeFromScope(elementId string, scope BaseElement) FlowElement {
	flowNode := FlowNode{}
	process := scope.(Process)
	if elementId != "" {
		element := process.GetFlowElement(elementId)
		flowNode = element.(FlowNode)
	}
	return flowNode
}
