package entity

import "github.com/lios/go-activiti/engine/impl/persistence/entity/data"

var (
	processDefinitionEntityManager        ProcessDefinitionEntityManagerImpl
	executionEntityManager                ExecutionEntityManagerImpl
	taskEntity                            TaskEntityManagerImpl
	identityLinkManager                   IdentityLinkEntityManagerImpl
	variableEntityManager                 VariableEntityManagerImpl
	historicTaskInstanceEntityManager     HistoricTaskInstanceEntityManagerImpl
	historicActivityInstanceEntityManager HistoricActivityInstanceEntityManagerImpl
	deploymentEntityManager               DeploymentEntityManagerImpl
	resourceEntityManager                 ResourceEntityManagerImpl
	historicIdentityLinkEntityManager     HistoricIdentityLinkEntityManagerImpl
)

func init() {
	processDefinitionEntityManager = ProcessDefinitionEntityManagerImpl{AbstractEntityManager{ProcessDefinitionEntityManagerImpl{}}}
	executionEntityManager = ExecutionEntityManagerImpl{AbstractEntityManager{ExecutionEntityManagerImpl{}}}
	taskEntity = TaskEntityManagerImpl{AbstractEntityManager{TaskEntityManagerImpl{}}}
	variableEntityManager = VariableEntityManagerImpl{AbstractEntityManager{VariableEntityManagerImpl{}}}
	identityLinkManager = IdentityLinkEntityManagerImpl{AbstractEntityManager{IdentityLinkEntityManagerImpl{}}}
	historicTaskInstanceEntityManager = HistoricTaskInstanceEntityManagerImpl{AbstractEntityManager: AbstractEntityManager{HistoricTaskInstanceEntityManagerImpl{}}}
	historicActivityInstanceEntityManager = HistoricActivityInstanceEntityManagerImpl{AbstractEntityManager: AbstractEntityManager{HistoricActivityInstanceEntityManagerImpl{}}}
	deploymentEntityManager = DeploymentEntityManagerImpl{AbstractEntityManager{DeploymentEntityManagerImpl{}}}
	resourceEntityManager = ResourceEntityManagerImpl{AbstractEntityManager{ResourceEntityManagerImpl{}}}
	historicIdentityLinkEntityManager = HistoricIdentityLinkEntityManagerImpl{AbstractEntityManager{HistoricIdentityLinkEntityManagerImpl{}}}

}

type EntityManager interface {
	GetDataManager() data.DataManagers

	Insert(interface{}) error

	GetById(id int64, data interface{}) interface{}

	Delete(entity Entity)
}

func GetTaskEntityManager() TaskEntityManagerImpl {
	return taskEntity
}
func GetProcessDefinitionEntityManager() ProcessDefinitionEntityManagerImpl {
	return processDefinitionEntityManager
}

func GetDeploymentEntityManager() DeploymentEntityManagerImpl {
	return deploymentEntityManager
}

func GetResourceDataManager() ResourceEntityManagerImpl {
	return resourceEntityManager
}
func GetExecutionEntityManager() ExecutionEntityManagerImpl {
	return executionEntityManager
}

func GetIdentityLinkManager() IdentityLinkEntityManagerImpl {
	return identityLinkManager
}

func GetVariableEntityManager() VariableEntityManagerImpl {
	return variableEntityManager
}

func GetHistoricTaskInstanceEntityManager() HistoricTaskInstanceEntityManagerImpl {
	return historicTaskInstanceEntityManager
}

func GetHistoricActivityInstanceEntityManager() HistoricActivityInstanceEntityManagerImpl {
	return historicActivityInstanceEntityManager
}

func GetHistoricIdentityLinkEntityManager() HistoricIdentityLinkEntityManagerImpl {
	return historicIdentityLinkEntityManager
}
