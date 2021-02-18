package model

import (
	"encoding/xml"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type SubProcess struct {
	delegate.BaseElement
	MessageName xml.Name `xml:"message"`
}
