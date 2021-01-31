package handler

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/bpmn/parse"
)

const name = "UserTaskParseHandler"

type UserTaskParseHandler struct {
}

func (userTaskParseHandler UserTaskParseHandler) GetHandledTypes() model.BaseElement {
	return model.UserTask{}
}

func (userTaskParseHandler UserTaskParseHandler) Parse(bpmnParse *parse.BpmnParse, flow model.BaseElement) {

}
