package model

import "encoding/xml"

//消息事件
type MessageEventDefinition struct {
	MessageEventDefinitionName xml.Name `xml:"messageEventDefinition"`
	MessageRef                 string   `xml:"messageRef,attr"`
}
