package model

import (
	"encoding/xml"
	"github.com/lios/go-activiti/engine/impl/bpmn"
)

type Message struct {
	bpmn.BaseElement
	MessageName xml.Name `xml:"message"`
}
