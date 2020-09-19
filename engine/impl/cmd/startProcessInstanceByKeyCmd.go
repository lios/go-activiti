package cmd

import (
	"github.com/lios/go-activiti/engine"
	"github.com/lios/go-activiti/engine/behavior"
	. "github.com/lios/go-activiti/engine/entityImpl"
	. "github.com/lios/go-activiti/engine/persistence"
	. "github.com/lios/go-activiti/model"
	"time"
)

type StartProcessInstanceByKeyCmd struct {
	ProcessDefinitionKey string
	Variables            map[string]interface{}
	BusinessKey          string
	TenantId             string
}

func (start StartProcessInstanceByKeyCmd) Execute(interceptor behavior.CommandContext) (interface{}, error) {
	defineManager := behavior.GetDefineManager()
	bytearries, err := defineManager.FindDeployedProcessDefinitionByKey(start.ProcessDefinitionKey)
	if err != nil {
		return nil, err
	}
	//解析xml数据
	process := behavior.GetBpmn(*bytearries[0])
	instance := ProcessInstance{}
	instance.BusinessKey = start.BusinessKey
	instance.TenantId = start.TenantId
	instance.StartTime = time.Now()
	instance.Key = process.Id
	instance.Name = process.Name
	instance.ProcessDefineId = bytearries[0].Id
	instance.DeploymentId = bytearries[0].DeploymentId
	//生成流程实例
	manager := ProcessInstanceManager{Instance: &instance}
	manager.CreateProcessInstance()
	//获取开始节点
	flowElement := process.InitialFlowElement
	element := flowElement.(engine.StartEvent)
	execution := ExecutionEntityImpl{ProcessInstanceId: instance.Id}
	execution.SetCurrentFlowElement(element)
	execution.SetProcessDefineId(bytearries[0].Id)
	execution.SetCurrentActivityId(element.GetId())
	//保存流程变量
	err = SetVariable(&execution, start.Variables)
	if err != nil {
		return nil, err
	}
	context, err := behavior.GetCommandContext()
	if err == nil {
		context.Agenda.PlanContinueProcessOperation(&execution)
	}
	return process, nil
}
