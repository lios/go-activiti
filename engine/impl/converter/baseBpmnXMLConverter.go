package converter

import (
	. "encoding/xml"
	"github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type BaseBpmnXMLConverter interface {
	GetXMLElementName() string

	convertToBpmnModel(decoder *Decoder, token StartElement, model BpmnModel, activeProcess Process)

	ConvertXMLToElement(decoder *Decoder, token StartElement, model BpmnModel, activeProcess Process) bpmn.BaseElement
}
