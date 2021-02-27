package cmd

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/interceptor"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
	"github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	"github.com/lios/go-activiti/engine/impl/utils"
	"github.com/lios/go-activiti/logger"
	. "github.com/lios/go-activiti/model"
	"time"
)

type StartProcessInstanceByKeyCmd struct {
	ProcessDefinitionKey string
	Variables            map[string]interface{}
	BusinessKey          string
	TenantId             string
}

func (start StartProcessInstanceByKeyCmd) Execute(command interceptor.CommandContext) (interface{}, error) {
	dataManager := entity.GetProcessDefinitionEntityManager()
	definitionEntity, err := dataManager.FindLatestProcessDefinitionByKey(start.ProcessDefinitionKey)
	if err != nil {
		logger.Error("FindLatestProcessDefinitionByKey err：", err)
	}
	processUtils := utils.ProcessDefinitionUtil{}
	process := processUtils.GetProcess(definitionEntity.GetId())
	instance := ProcessInstance{}
	instance.BusinessKey = start.BusinessKey
	instance.TenantId = start.TenantId
	instance.StartTime = time.Now()
	instance.Key = process.Id
	instance.Name = process.Name
	instance.ProcessDefineId = definitionEntity.GetId()
	instance.DeploymentId = definitionEntity.GetDeploymentId()
	//生成流程实例
	manager := data.ProcessInstanceDataManager{Instance: &instance}
	manager.CreateProcessInstance()
	//获取开始节点
	flowElement := process.InitialFlowElement
	element := flowElement.(*model.StartEvent)
	execution := entity.ExecutionEntityImpl{ProcessInstanceId: instance.Id}
	execution.SetCurrentFlowElement(element)
	//execution.SetProcessDefineId(bytearries[0].Id)
	execution.SetCurrentActivityId(element.GetId())
	//保存流程变量
	err = execution.SetVariable(&execution, start.Variables)
	if err != nil {
		return nil, err
	}
	context, err := interceptor.GetCommandContext()
	if err == nil {
		context.Agenda.PlanContinueProcessOperation(&execution)
	}
	return process, nil
}
