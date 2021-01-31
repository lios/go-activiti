package parse

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type BpmnParseHandler interface {
	GetHandledTypes() model.BaseElement

	Parse(bpmnParse *BpmnParse, flow model.BaseElement)
}
