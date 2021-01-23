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
