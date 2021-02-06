package entity

type ProcessInstanceEntity interface {
	SetName(name string)

	SetKey(key string)

	GetProcessDefinitionId() string

	SetProcessDefinitionId(processDefinitionId string)
}
