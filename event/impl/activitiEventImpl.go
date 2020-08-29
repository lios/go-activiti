package impl

import . "github.com/lios/go-activiti/event"

type ActivitiEventImpl struct {
	EventType           ActivitiEventType
	ExecutionId         string
	ProcessInstanceId   string
	ProcessDefinitionId string
}

func (activitiEvent ActivitiEventImpl) GetType() ActivitiEventType {
	return activitiEvent.EventType
}

func (activitiEvent ActivitiEventImpl) SetType(eventType ActivitiEventType) {
	activitiEvent.EventType = eventType
}
func (activitiEvent ActivitiEventImpl) GetProcessDefinitionId() string {
	return activitiEvent.ProcessDefinitionId
}

func (activitiEvent ActivitiEventImpl) SetProcessDefinitionId(processDefinitionId string) {
	activitiEvent.ProcessDefinitionId = processDefinitionId
}

func (activitiEvent ActivitiEventImpl) GetProcessInstanceId() string {
	return activitiEvent.ProcessInstanceId
}

func (activitiEvent ActivitiEventImpl) SetProcessInstanceId(processInstanceId string) {
	activitiEvent.ProcessInstanceId = processInstanceId
}
