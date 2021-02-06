package model

import "github.com/lios/go-activiti/engine/impl/bpmn"

type FlowElementsContainer interface {
	AddFlowElement(element bpmn.FlowElement)
}
