package model

import "encoding/xml"

type EndEvent struct {
	*Flow
	EndEventName xml.Name `xml:"endEvent"`
}
