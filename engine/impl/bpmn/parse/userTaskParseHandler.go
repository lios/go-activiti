package parse

import (
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/delegate"
)

type UserTaskParseHandler struct {
	AbstractActivityBpmnParseHandler
}

func (userTaskParseHandler UserTaskParseHandler) GetHandledType() string {
	return UserTask{}.GetType()
}

func (userTaskParseHandler UserTaskParseHandler) ExecuteParse(bpmnParse BpmnParse, baseElement delegate.BaseElement) {
	userTask := baseElement.(*UserTask)
	userTask.SetBehavior(bpmnParse.ActivityBehaviorFactory.CreateUserTaskActivityBehavior(*userTask))
}
