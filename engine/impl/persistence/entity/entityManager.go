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
)

type EntityManager interface {
	Entity
	GetDataManager() data.DataManagers

	Insert(interface{}) error

	GetById(id int64) Entity

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

func GetIdentityLinkManager() IdentityLinkEntityManager {
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
