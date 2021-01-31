package model

import "encoding/xml"

type UserTask struct {
	*Flow
	UserTaskName      xml.Name                         `xml:"userTask"`
	Assignee          string                           `xml:"assignee,attr"`
	CandidateUsers    []string                         `xml:"candidateUsers,attr"`
	CandidateGroups   []string                         `xml:"candidateGroups,attr"`
	ExtensionElements ExtensionElement                 `xml:"extensionElements"`
	MultiInstance     MultiInstanceLoopCharacteristics `xml:"multiInstanceLoopCharacteristics"`
}
