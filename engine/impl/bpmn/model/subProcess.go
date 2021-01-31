package model

import "encoding/xml"

type SubProcess struct {
	BaseElement
	MessageName xml.Name `xml:"message"`
}
