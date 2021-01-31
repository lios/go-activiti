package agenda

import (
	"container/list"
	"github.com/lios/go-activiti/engine/impl/invocation"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type DefaultActivitiEngineAgenda struct {
	Operations list.List
}

//判断是否为空
func (agenda *DefaultActivitiEngineAgenda) IsEmpty() bool {
	return agenda.Operations.Len() == 0
}

//设置后续操作
func (agenda *DefaultActivitiEngineAgenda) PlanOperation(operation invocation.Operation) {
	agenda.Operations.PushFront(operation)
}

func (agenda *DefaultActivitiEngineAgenda) GetNextOperation() invocation.Operation {
	value := agenda.Operations.Front()
	agenda.Operations.Remove(value)
	return value.Value.(invocation.Operation)
}

//获取下一步操作
func (agenda *DefaultActivitiEngineAgenda) getNextOperation() invocation.Operation {
	return agenda.Operations.Front().Value.(invocation.Operation)
}

//连线继续执行
func (agenda *DefaultActivitiEngineAgenda) PlanContinueProcessOperation(execution entity.ExecutionEntity) {
	agenda.PlanOperation(&ContinueProcessOperation{AbstractOperation{Execution: execution}})
}

//任务出口执行
func (agenda *DefaultActivitiEngineAgenda) PlanTriggerExecutionOperation(execution entity.ExecutionEntity) {
	agenda.PlanOperation(&TriggerExecutionOperation{AbstractOperation{Execution: execution}})
}

//连线出口设置
func (agenda *DefaultActivitiEngineAgenda) PlanTakeOutgoingSequenceFlowsOperation(execution entity.ExecutionEntity, valuateConditions bool) {
	agenda.PlanOperation(&TakeOutgoingSequenceFlowsOperation{AbstractOperation: AbstractOperation{Execution: execution}, EvaluateConditions: valuateConditions})
}

//任务结束
func (agenda *DefaultActivitiEngineAgenda) PlanEndExecutionOperation(execution entity.ExecutionEntity) {
	agenda.PlanOperation(&EndExecutionOperation{AbstractOperation: AbstractOperation{Execution: execution}})
}
