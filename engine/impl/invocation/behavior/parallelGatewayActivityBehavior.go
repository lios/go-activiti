package behavior

import (
	"github.com/lios/go-activiti/engine/impl/delegate"
	"github.com/lios/go-activiti/engine/impl/interceptor"
	"github.com/lios/go-activiti/engine/impl/manager"
	"github.com/lios/go-activiti/engine/impl/utils"
)

type ParallelGatewayActivityBehavior struct {
}

//包容网关
func (parallel ParallelGatewayActivityBehavior) Execute(execution delegate.DelegateExecution) error {
	return parallel.Leave(execution)
}

//执行逻辑：获取当前所有执行的节点，判断是否可达当前网关可以停止执行，等待完成
func (parallel ParallelGatewayActivityBehavior) Leave(execution delegate.DelegateExecution) error {
	processInstanceId := execution.GetProcessInstanceId()
	taskManager := manager.GetDataManager().TaskDataManager
	//查询当前执行节点
	tasks, errS := taskManager.FindByProcessInstanceId(processInstanceId)
	var oneExecutionCanReachGateway = false
	if errS != nil {
		for _, task := range tasks {
			if task.TaskDefineKey != execution.GetCurrentActivityId() {
				//判断是否可以继续执行
				oneExecutionCanReachGateway = utils.IsReachable(execution.GetProcessDefineId(), task.TaskDefineKey, execution.GetCurrentActivityId())
			} else {
				oneExecutionCanReachGateway = true
			}
		}
	}
	if !oneExecutionCanReachGateway {
		//执行出口逻辑，设置条件判断
		interceptor.GetAgenda().Agenda.PlanTakeOutgoingSequenceFlowsOperation(execution, true)
	}
	return nil
}
