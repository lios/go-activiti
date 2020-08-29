package engine

type ActivitiEngineAgenda interface {
	IsEmpty() bool

	PlanOperation(operation Operation)

	GetNextOperation() Operation

	PlanContinueProcessOperation(execution ExecutionEntity)

	//planContinueProcessSynchronousOperation(execution ExecutionEntity)
	//
	//planContinueProcessInCompensation(execution ExecutionEntity)
	//
	//planContinueMultiInstanceOperation(execution ExecutionEntity)

	PlanTakeOutgoingSequenceFlowsOperation(execution ExecutionEntity, evaluateConditions bool)

	PlanEndExecutionOperation(execution ExecutionEntity)

	PlanTriggerExecutionOperation(execution ExecutionEntity)
	//
	//planDestroyScopeOperation(execution ExecutionEntity)
	//
	//planExecuteInactiveBehaviorsOperation()

}
