package parse

import (
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type BpmnParseHandler interface {
	GetHandledTypes() []string

	Parse(bpmnParse *BpmnParse, flow delegate.BaseElement)
}
