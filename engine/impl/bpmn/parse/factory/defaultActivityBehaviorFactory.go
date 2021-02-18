package factory

import (
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	. "github.com/lios/go-activiti/engine/impl/invocation/behavior"
)

type DefaultActivityBehaviorFactory struct {
}

func (defaultActivityBehaviorFactory DefaultActivityBehaviorFactory) CreateUserTaskActivityBehavior(userTask UserTask, key string) UserTaskActivityBehavior {
	return UserTaskActivityBehavior{UserTask: userTask, ProcessKey: key}
}

func (defaultActivityBehaviorFactory DefaultActivityBehaviorFactory) CreateAutoUserTaskActivityBehavior(userTask UserTask, key string) UserAutoTaskActivityBehavior {
	return UserAutoTaskActivityBehavior{UserTask: userTask, ProcessKey: key}
}

func (defaultActivityBehaviorFactory DefaultActivityBehaviorFactory) CreateExclusiveGatewayActivityBehavior(exclusiveGateway ExclusiveGateway) ExclusiveGatewayActivityBehavior {
	return ExclusiveGatewayActivityBehavior{}
}

func (defaultActivityBehaviorFactory DefaultActivityBehaviorFactory) CreateInclusiveGatewayActivityBehavior(inclusiveGateway InclusiveGateway) InclusiveGatewayActivityBehavior {
	return InclusiveGatewayActivityBehavior{}
}
