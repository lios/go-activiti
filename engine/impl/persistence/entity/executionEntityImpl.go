package entity

import (
	"github.com/lios/go-activiti/engine/impl/delegate"
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	. "github.com/lios/go-activiti/engine/variable"
)

type ExecutionEntityImpl struct {
	AbstractEntity
	AbstractEntityManager
	VariableScopeImpl
	BusinessKey        string
	CurrentFlowElement delegate.FlowElement
	DeploymentId       int
	ProcessInstanceId  int64
	ProcessDefineId    int64
	CurrentActivityId  string
}

func (execution *ExecutionEntityImpl) SetBusinessKey(businessKey string) {
	execution.BusinessKey = businessKey
}

func (execution ExecutionEntityImpl) GetCurrentFlowElement() delegate.FlowElement {
	return execution.CurrentFlowElement
}

func (execution *ExecutionEntityImpl) SetCurrentFlowElement(flow delegate.FlowElement) {
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

func (execution ExecutionEntityImpl) GetProcessDefineId() int64 {
	return execution.ProcessDefineId
}

func (execution *ExecutionEntityImpl) SetProcessDefineId(processDefineId int64) {
	execution.ProcessDefineId = processDefineId
}

func (execution ExecutionEntityImpl) GetCurrentActivityId() string {
	return execution.CurrentActivityId
}

func (execution *ExecutionEntityImpl) SetCurrentActivityId(currentActivityId string) {
	execution.CurrentActivityId = currentActivityId
}

func (execution ExecutionEntityImpl) GetTaskId() int64 {
	return -1
}

func (execution *ExecutionEntityImpl) SetTaskId(taskId int64) {

}
func (execution ExecutionEntityImpl) GetProcessVariable() map[string]interface{} {
	return execution.GetVariable()
}

func (execution ExecutionEntityImpl) GetVariable() map[string]interface{} {
	dateManager := GetVariableEntityManager().GetDataManager()
	variableManager := dateManager.(VariableDataManager)
	variables, err := variableManager.SelectByProcessInstanceId(execution.GetProcessInstanceId())
	if err == nil {
		return execution.HandleVariable(variables)
	}
	return nil
}

func (execution ExecutionEntityImpl) HandleVariable(variables []Variable) map[string]interface{} {
	variableManager := GetVariableManager()
	variableTypes := variableManager.VariableTypes
	var variableMap = make(map[string]interface{}, 0)
	for _, variable := range variables {
		variableType := variableTypes.GetVariableType(variable.Type)
		value := variableType.GetValue(&variable)
		variableMap[variable.Name] = value
	}
	return variableMap
}

//保存流程变量
//func (execution ExecutionEntityImpl) SetVariable(variables map[string]interface{}) error {
//variableManager :=  GetVariableManager()
//variableTypes := variableManager.VariableTypes
//variableDataManager := execution.GetVariableEntityManager()
//if variables != nil && len(variables) > 0 {
//	for k, v := range variables {
//		kind := reflect.TypeOf(v).Kind()
//		variableType := variableTypes.GetVariableType(kind.String())
//		if variableType == nil {
//			return errs.ProcessError{Code: "1001", Msg: "no type"}
//		}
//		variable := variableDataManager.Create(k, variableType, v)
//		//存在更新
//		specificVariable, e := execution.GetSpecificVariable(k)
//		if e != nil {
//			variable.Version = specificVariable.Version + 1
//		}
//		execution.SetScope(variable)
//		variableDataManager.Variable = variable
//		err := variableDataManager.Insert()
//		if err != nil {
//			return err
//		}
//	}
//}
//	return nil
//}
func (execution ExecutionEntityImpl) GetSourceActivityExecution() ExecutionEntity {
	return &execution
}
func (execution ExecutionEntityImpl) GetSpecificVariable(variableName string) (Variable, error) {
	manager := GetVariableEntityManager().GetDataManager()
	variableDataManager := manager.(VariableDataManager)
	return variableDataManager.SelectProcessInstanceId(variableName, execution.ProcessInstanceId)
	return Variable{}, nil
}

func (execution ExecutionEntityImpl) SetScope(variable *Variable) {
	variable.ProcessInstanceId = execution.ProcessInstanceId
}
