package model

import "encoding/xml"

type Message struct {
	BaseElement
	MessageName xml.Name `xml:"message"`
}
