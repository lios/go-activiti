package model

import (
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type FlowElementsContainer interface {
	AddFlowElement(element delegate.FlowElement)
}
