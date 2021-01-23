package manager

import . "github.com/lios/go-activiti/engine/persistence"

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
