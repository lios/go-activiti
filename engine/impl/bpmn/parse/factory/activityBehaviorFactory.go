package factory

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/invocation/behavior"
)

type ActivityBehaviorFactory interface {
	CreateUserTaskActivityBehavior(userTask model.UserTask) behavior.UserTaskActivityBehavior
}
