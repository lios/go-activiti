package behavior

import (
	"github.com/lios/go-activiti/engine"
	. "github.com/lios/go-activiti/engine/persistence"
)

type CommandContext struct {
	Command                    Command
	Agenda                     engine.ActivitiEngineAgenda
	ProcessEngineConfiguration ProcessEngineConfiguration
}

func GetProcessInstanceManager() ProcessInstanceManager {
	return ProcessInstanceManager{}
}

func GetTaskManager() TaskManager {
	return TaskManager{}
}

func GetDefineManager() DefineManager {
	return DefineManager{}
}
func GetVariableManager() VariableManager {
	return VariableManager{}
}

func GetIdentityLinkManager() IdentityLinkManager {
	return IdentityLinkManager{}
}

func GetHistoricActinstManager() HistoricActinstManager {
	return HistoricActinstManager{}
}

func GetHistoricTaskManager() HistoricTaskManager {
	return HistoricTaskManager{}
}

func GetHistoricProcessManager() HistoricProcessManager {
	return HistoricProcessManager{}
}
