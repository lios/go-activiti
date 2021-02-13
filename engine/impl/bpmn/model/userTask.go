package model

import "encoding/xml"

type UserTask struct {
	FlowNode
	UserTaskName      xml.Name                         `xml:"userTask"`
	Assignee          string                           `xml:"assignee,attr"`
	CandidateUsers    []string                         `xml:"candidateUsers,attr"`
	CandidateGroups   []string                         `xml:"candidateGroups,attr"`
	ExtensionElements ExtensionElement                 `xml:"extensionElements"`
	MultiInstance     MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics"`
}

func (user UserTask) GetType() string {
	return "UserTask"
}
