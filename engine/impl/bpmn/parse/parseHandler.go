package parse

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
)

type ParseHandler interface {
	GetHandledType() string

	ExecuteParse(bpmnParse BpmnParse, element bpmn.BaseElement)
}
