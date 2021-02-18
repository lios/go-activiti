package parse

import (
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type ParseHandler interface {
	GetHandledType() string

	ExecuteParse(bpmnParse BpmnParse, element delegate.BaseElement)
}
