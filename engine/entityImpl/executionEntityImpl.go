package entityImpl

import (
	"github.com/lios/go-activiti/engine"
	. "github.com/lios/go-activiti/engine/behavior"
	. "github.com/lios/go-activiti/engine/manager"
	. "github.com/lios/go-activiti/engine/persistence"
	. "github.com/lios/go-activiti/engine/variable"
	"github.com/lios/go-activiti/errs"
	"reflect"
)

type ExecutionEntityImpl struct {
	BusinessKey        string
	CurrentFlowElement engine.FlowElement
	DeploymentId       int
	ProcessInstanceId  int64
	ProcessDefineId    int64
	CurrentActivityId  string
}

func (execution *ExecutionEntityImpl) SetBusinessKey(businessKey string) {
	execution.BusinessKey = businessKey
}

func (execution ExecutionEntityImpl) GetCurrentFlowElement() engine.FlowElement {
	return execution.CurrentFlowElement
}

func (execution *ExecutionEntityImpl) SetCurrentFlowElement(flow engine.FlowElement) {
	execution.CurrentFlowElement = flow
	execution.CurrentActivityId = flow.GetId()
}

func (execution ExecutionEntityImpl) GetDeploymentId() int {
	return execution.DeploymentId
}

func (execution *ExecutionEntityImpl) SetDeploymentId(deploymentId int) {
	execution.DeploymentId = deploymentId
}

func (execution ExecutionEntityImpl) GetProcessInstanceId() int64 {
	return execution.ProcessInstanceId
}

func (execution *ExecutionEntityImpl) SetProcessInstanceId(processInstanceId int64) {
	execution.ProcessInstanceId = processInstanceId
}

func (execution *ExecutionEntityImpl) GetProcessDefineId() int64 {
	return execution.ProcessDefineId
}

func (execution *ExecutionEntityImpl) SetProcessDefineId(processDefineId int64) {
	execution.ProcessDefineId = processDefineId
}

func (execution *ExecutionEntityImpl) GetCurrentActivityId() string {
	return execution.CurrentActivityId
}

func (execution *ExecutionEntityImpl) SetCurrentActivityId(currentActivityId string) {
	execution.CurrentActivityId = currentActivityId
}

func (execution *ExecutionEntityImpl) GetTaskId() int64 {
	return -1
}

func (execution *ExecutionEntityImpl) SetTaskId(taskId int64) {

}
func (execution *ExecutionEntityImpl) GetProcessVariable() map[string]interface{} {
	return execution.GetVariable()
}

func (execution *ExecutionEntityImpl) GetVariable() map[string]interface{} {
	variableManager := GetVariableManager()
	variables, err := variableManager.SelectByProcessInstanceId(execution.GetProcessInstanceId())
	if err == nil {
		return execution.HandleVariable(variables)
	}
	return nil
}

func (execution *ExecutionEntityImpl) HandleVariable(variables []Variable) map[string]interface{} {
	engineConfiguration := GetProcessEngineConfiguration()
	variableTypes := engineConfiguration.VariableTypes
	var variableMap = make(map[string]interface{}, 0)
	for _, variable := range variables {
		variableType := variableTypes.GetVariableType(variable.Type)
		value := variableType.GetValue(&variable)
		variableMap[variable.Name] = value
	}
	return variableMap
}

//保存流程变量
func SetVariable(execution engine.ExecutionEntity, variables map[string]interface{}) error {
	engineConfiguration := GetProcessEngineConfiguration()
	variableTypes := engineConfiguration.VariableTypes
	variableManager := GetVariableManager()
	if variables != nil && len(variables) > 0 {
		for k, v := range variables {
			kind := reflect.TypeOf(v).Kind()
			variableType := variableTypes.GetVariableType(kind.String())
			if variableType == nil {
				return errs.ProcessError{Code: "1001", Msg: "no type"}
			}
			variable := variableManager.Create(k, variableType, v)
			//存在更新
			specificVariable, e := execution.GetSpecificVariable(k)
			if e != nil {
				variable.Version = specificVariable.Version + 1
			}
			execution.SetScope(variable)
			variableManager.Variable = variable
			err := variableManager.Insert()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (execution *ExecutionEntityImpl) GetSpecificVariable(variableName string) (Variable, error) {
	variableManager := VariableManager{}
	return variableManager.SelectProcessInstanceId(variableName, execution.ProcessInstanceId)
}

func (execution *ExecutionEntityImpl) SetScope(variable *Variable) {
	variable.ProcessInstanceId = execution.ProcessInstanceId
}
