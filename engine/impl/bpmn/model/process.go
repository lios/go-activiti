package model

import (
	"encoding/xml"
	. "github.com/lios/go-activiti/engine/impl/bpmn"
)

type Process struct {
	FlowElement
	ProcessName            xml.Name                 `xml:"process"`
	Id                     string                   `xml:"id,attr"`
	Name                   string                   `xml:"name,attr"`
	Documentation          string                   `xml:"documentation"`
	IsExecutable           string                   `xml:"isExecutable,attr"`
	StartEvent             []StartEvent             `xml:"startEvent"`
	EndEvent               []EndEvent               `xml:"endEvent"`
	UserTask               []UserTask               `xml:"userTask"`
	SequenceFlow           []SequenceFlow           `xml:"sequenceFlow"`
	ExclusiveGateway       []ExclusiveGateway       `xml:"exclusiveGateway"`
	InclusiveGateway       []InclusiveGateway       `xml:"inclusiveGateway"`
	ParallelGateway        []ParallelGateway        `xml:"parallelGateway"`
	BoundaryEvent          []BoundaryEvent          `xml:"boundaryEvent"`
	IntermediateCatchEvent []IntermediateCatchEvent `xml:"intermediateCatchEvent"`
	SubProcess             []SubProcess             `xml:"subProcess"`
	FlowElementList        []FlowElement
	InitialFlowElement     FlowElement
	FlowElementMap         map[string]FlowElement
}

func (process Process) GetFlowElement(flowElementId string) FlowElement {
	return process.FlowElementMap[flowElementId]
}

func (process Process) AddFlowElement(element FlowElement) {
	process.FlowElementList = append(process.FlowElementList, element)
	process.FlowElementMap[element.GetId()] = element
}
