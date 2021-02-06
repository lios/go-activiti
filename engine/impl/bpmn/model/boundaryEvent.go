package model

import "encoding/xml"

type BoundaryEvent struct {
	*FlowNode
	BoundaryEventName    xml.Name             `xml:"boundaryEvent"`
	AttachedToRef        string               `xml:"attachedToRef,attr"`
	CancelActivity       string               `xml:"cancelActivity,attr"`
	TimerEventDefinition TimerEventDefinition `xml:"timerEventDefinition"`
}
