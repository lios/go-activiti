package handler

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/parse"
)

type ParseHandler interface {
	GetHandledType() bpmn.BaseElement

	ExecuteParse(bpmnParse BpmnParse, element bpmn.BaseElement)
}
