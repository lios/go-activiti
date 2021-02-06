package model

import "encoding/xml"

type EndEvent struct {
	*FlowNode
	EndEventName xml.Name `xml:"endEvent"`
}
