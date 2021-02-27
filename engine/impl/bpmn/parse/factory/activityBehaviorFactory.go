package factory

import (
	. "github.com/lios/go-activiti/engine/impl/bpmn/model"
	. "github.com/lios/go-activiti/engine/impl/invocation/behavior"
)

type ActivityBehaviorFactory interface {
	CreateUserTaskActivityBehavior(userTask UserTask, key string) UserTaskActivityBehavior

	CreateAutoUserTaskActivityBehavior(userTask UserTask, key string) UserAutoTaskActivityBehavior

	CreateExclusiveGatewayActivityBehavior(exclusiveGateway ExclusiveGateway) ExclusiveGatewayActivityBehavior

	CreateInclusiveGatewayActivityBehavior(inclusiveGateway InclusiveGateway) InclusiveGatewayActivityBehavior

	CreateParallelGatewayActivityBehavior(inclusiveGateway ParallelGateway) ParallelGatewayActivityBehavior
}
