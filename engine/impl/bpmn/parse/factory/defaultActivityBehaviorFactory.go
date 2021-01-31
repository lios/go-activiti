package factory

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	. "github.com/lios/go-activiti/engine/impl/invocation/behavior"
)

type DefaultActivityBehaviorFactory struct {
}

func (defaultActivityBehaviorFactory DefaultActivityBehaviorFactory) CreateUserTaskActivityBehavior(userTask model.UserTask) UserTaskActivityBehavior {
	return UserTaskActivityBehavior{UserTask: userTask}
}
