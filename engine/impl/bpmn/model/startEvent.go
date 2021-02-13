package model

import "encoding/xml"

type StartEvent struct {
	FlowNode
	StartEventName xml.Name `xml:"startEvent"`
	Initiator      string   `xml:"initiator,attr"`
	FormKey        string   `xml:"formKey,attr"`
}

func (start StartEvent) GetType() string {
	return "StartEvent"
}
