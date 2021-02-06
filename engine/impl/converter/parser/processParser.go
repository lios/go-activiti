package parser

import (
	. "encoding/xml"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type ProcessParser struct {
}

func (ProcessParser ProcessParser) Parse(decoder *Decoder, token StartElement, model BpmnModel) Process {
	process := &Process{}
	decoder.DecodeElement(process, &token)
	model.AddProcess(process)
	return *process
}
