package behavior

import (
	"github.com/lios/go-activiti/engine"
	. "github.com/lios/go-activiti/engine/manager"
	"github.com/lios/go-activiti/engine/utils"
)

type InclusiveGatewayActivityBehavior struct {
}

//包容网关
func (exclusive InclusiveGatewayActivityBehavior) Execute(execution engine.ExecutionEntity) error {
	return exclusive.Leave(execution)
}

//执行逻辑：获取当前所有执行的节点，判断是否可达当前网关可以停止执行，等待完成
func (exclusive InclusiveGatewayActivityBehavior) Leave(execution engine.ExecutionEntity) error {
	processInstanceId := execution.GetProcessInstanceId()
	taskManager := GetTaskManager()
	//查询当前执行节点
	tasks, errS := taskManager.FindByProcessInstanceId(processInstanceId)
	var oneExecutionCanReachGateway = false
	if errS != nil {
		bytearry, err := GetDefineManager().GetBytearry(execution.GetProcessDefineId())
		if err != nil {
			return err
		}
		process := GetBpmn(bytearry)
		for _, task := range tasks {
			if task.TaskDefineKey != execution.GetCurrentActivityId() {
				//判断是否可以继续执行
				oneExecutionCanReachGateway = utils.IsReachable(process, task.TaskDefineKey, execution.GetCurrentActivityId())
			} else {
				oneExecutionCanReachGateway = true
			}
		}
	}
	if !oneExecutionCanReachGateway {
		//执行出口逻辑，设置条件判断
		GetAgenda().PlanTakeOutgoingSequenceFlowsOperation(execution, true)
	}
	return nil
}
