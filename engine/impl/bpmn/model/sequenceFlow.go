package model

import "encoding/xml"

type SequenceFlow struct {
	*FlowNode
	SequenceFlowName    xml.Name `xml:"sequenceFlow"`
	Id                  string   `xml:"id,attr"`
	SourceRef           string   `xml:"sourceRef,attr"`
	TargetRef           string   `xml:"targetRef,attr"`
	ConditionExpression string   `xml:"conditionExpression"`
}
