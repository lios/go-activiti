package handler

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/bpmn/parse"
)

type UserTaskParseHandler struct {
	AbstractActivityBpmnParseHandler
}

func (userTaskParseHandler UserTaskParseHandler) GetHandledType() bpmn.BaseElement {
	return UserTask{}
}

func (userTaskParseHandler UserTaskParseHandler) ExecuteParse(bpmnParse parse.BpmnParse, baseElement bpmn.BaseElement) {
	userTask := baseElement.(UserTask)
	userTask.SetBehavior(bpmnParse.ActivityBehaviorFactory.CreateUserTaskActivityBehavior(userTask))
}
