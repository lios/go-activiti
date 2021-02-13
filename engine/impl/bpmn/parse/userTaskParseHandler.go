package parse

import (
	"github.com/lios/go-activiti/engine/impl/bpmn"
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
)

type UserTaskParseHandler struct {
	AbstractActivityBpmnParseHandler
}

func (userTaskParseHandler UserTaskParseHandler) GetHandledType() string {
	return UserTask{}.GetType()
}

func (userTaskParseHandler UserTaskParseHandler) ExecuteParse(bpmnParse BpmnParse, baseElement bpmn.BaseElement) {
	userTask := baseElement.(*UserTask)
	userTask.SetBehavior(bpmnParse.ActivityBehaviorFactory.CreateUserTaskActivityBehavior(*userTask))
}
