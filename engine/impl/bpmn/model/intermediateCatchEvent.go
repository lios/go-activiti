package model

import "encoding/xml"

//中间抛出事件
type IntermediateCatchEvent struct {
	*Flow
	IntermediateCatchEventName xml.Name               `xml:"intermediateCatchEvent"`
	MessageEventDefinition     MessageEventDefinition `xml:"messageEventDefinition"`
}
