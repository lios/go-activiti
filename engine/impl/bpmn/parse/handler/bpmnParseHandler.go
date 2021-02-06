package handler

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	"github.com/lios/go-activiti/engine/impl/bpmn/parse"
)

type BpmnParseHandler interface {
	GetHandledTypes() []bpmn.BaseElement

	Parse(bpmnParse *parse.BpmnParse, flow bpmn.BaseElement)
}
