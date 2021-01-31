package model

import "encoding/xml"

type TimerEventDefinition struct {
	TimerEventDefinitionName xml.Name `xml:"timerEventDefinition"`
	TimeDuration             string   `xml:"timeDuration"`
}
