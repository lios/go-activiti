package behavior

import (
	"container/list"
	"github.com/lios/go-activiti/engine"
)

type DefaultActivitiEngineAgenda struct {
	Operations list.List
}

//判断是否为空
func (agenda *DefaultActivitiEngineAgenda) IsEmpty() bool {
	return agenda.Operations.Len() == 0
}

//设置后续操作
func (agenda *DefaultActivitiEngineAgenda) PlanOperation(operation engine.Operation) {
	agenda.Operations.PushFront(operation)
}

func (agenda *DefaultActivitiEngineAgenda) GetNextOperation() engine.Operation {
	value := agenda.Operations.Front()
	agenda.Operations.Remove(value)
	return value.Value.(engine.Operation)
}

//获取下一步操作
func (agenda *DefaultActivitiEngineAgenda) getNextOperation() engine.Operation {
	return agenda.Operations.Front().Value.(engine.Operation)
}

//连线继续执行
func (agenda *DefaultActivitiEngineAgenda) PlanContinueProcessOperation(execution engine.ExecutionEntity) {
	agenda.PlanOperation(&ContinueProcessOperation{AbstractOperation{Execution: execution}})
}

//任务出口执行
func (agenda *DefaultActivitiEngineAgenda) PlanTriggerExecutionOperation(execution engine.ExecutionEntity) {
	agenda.PlanOperation(&TriggerExecutionOperation{AbstractOperation{Execution: execution}})
}

//连线出口设置
func (agenda *DefaultActivitiEngineAgenda) PlanTakeOutgoingSequenceFlowsOperation(execution engine.ExecutionEntity, valuateConditions bool) {
	agenda.PlanOperation(&TakeOutgoingSequenceFlowsOperation{AbstractOperation: AbstractOperation{Execution: execution}, EvaluateConditions: valuateConditions})
}

//任务结束
func (agenda *DefaultActivitiEngineAgenda) PlanEndExecutionOperation(execution engine.ExecutionEntity) {
	agenda.PlanOperation(&EndExecutionOperation{AbstractOperation: AbstractOperation{Execution: execution}})
}
