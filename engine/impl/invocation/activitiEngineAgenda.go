package invocation

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type ActivitiEngineAgenda interface {
	IsEmpty() bool

	PlanOperation(operation Operation)

	GetNextOperation() Operation

	PlanContinueProcessOperation(execution entity.ExecutionEntity)

	//planContinueProcessSynchronousOperation(execution ExecutionEntity)
	//
	//planContinueProcessInCompensation(execution ExecutionEntity)
	//
	//planContinueMultiInstanceOperation(execution ExecutionEntity)

	PlanTakeOutgoingSequenceFlowsOperation(execution entity.ExecutionEntity, evaluateConditions bool)

	PlanEndExecutionOperation(execution entity.ExecutionEntity)

	PlanTriggerExecutionOperation(execution entity.ExecutionEntity)
	//
	//planDestroyScopeOperation(execution ExecutionEntity)
	//
	//planExecuteInactiveBehaviorsOperation()

}
