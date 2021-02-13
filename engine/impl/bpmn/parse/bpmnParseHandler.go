package parse

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
)

type BpmnParseHandler interface {
	GetHandledTypes() []string

	Parse(bpmnParse *BpmnParse, flow bpmn.BaseElement)
}
