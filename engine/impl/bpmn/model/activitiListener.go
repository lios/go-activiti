package model

import "encoding/xml"

type ActivitiListener struct {
	TaskListenerName xml.Name `xml:"taskListener"`
	EventType        string   `xml:"event,attr"`
}
