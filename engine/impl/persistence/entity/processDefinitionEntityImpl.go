package entity

type ProcessDefinitionEntityImpl struct {
	AbstractEntity
	ExecutionEntityImpl
	Name         string
	Description  string
	Key          string
	Version      int
	Category     string
	DeploymentId int64
	ResourceName string
}

func (processDefinitionEntityManager *ProcessDefinitionEntityImpl) SetKey(key string) {
	processDefinitionEntityManager.Key = key
}

func (processDefinitionEntityManager *ProcessDefinitionEntityImpl) SetName(name string) {
	processDefinitionEntityManager.Name = name
}

func (processDefinitionEntityManager *ProcessDefinitionEntityImpl) SetDescription(description string) {
	processDefinitionEntityManager.Description = description
}

func (processDefinitionEntityManager *ProcessDefinitionEntityImpl) SetDeploymentId(deploymentId int64) {
	processDefinitionEntityManager.DeploymentId = deploymentId
}

func (processDefinitionEntityManager *ProcessDefinitionEntityImpl) SetVersion(version int) {
	processDefinitionEntityManager.Version = version
}

func (processDefinitionEntityManager *ProcessDefinitionEntityImpl) SetResourceName(resourceName string) {
	processDefinitionEntityManager.ResourceName = resourceName
}

func (processDefinitionEntityManager *ProcessDefinitionEntityImpl) GetKey() string {
	return processDefinitionEntityManager.Key
}

func (processDefinitionEntityManager *ProcessDefinitionEntityImpl) GetName() string {
	return processDefinitionEntityManager.Key
}

func (processDefinitionEntityManager *ProcessDefinitionEntityImpl) GetDescription() string {
	return processDefinitionEntityManager.Description
}
func (processDefinitionEntityManager *ProcessDefinitionEntityImpl) GetDeploymentId() int64 {
	return processDefinitionEntityManager.DeploymentId
}

func (processDefinitionEntityManager *ProcessDefinitionEntityImpl) GetVersion() int {
	return processDefinitionEntityManager.Version
}

func (processDefinitionEntityManager *ProcessDefinitionEntityImpl) GetResourceName() string {
	return processDefinitionEntityManager.ResourceName
}
