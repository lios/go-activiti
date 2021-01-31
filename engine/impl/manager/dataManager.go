package manager

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
)

var dataManager DataManager

type DataManager struct {
	DefineDataManager               *DefineDataManager
	HistoricActinstDataManager      *HistoricActinstDataManager
	HistoricIdentityLinkDataManager *HistoricIdentityLinkDataManager
	HistoricProcessDataManager      *HistoricProcessDataManager
	HistoricTaskDataManager         *HistoricTaskDataManager
	HistoricVariableDataManager     *HistoricVariableDataManager
	IdentityLinkDataManager         *IdentityLinkDataManager
	ProcessInstanceDataManager      *ProcessInstanceDataManager
	TaskDataManager                 *TaskDataManager
	VariableDataManager             *VariableDataManager
}

func init() {
	dataManager = DataManager{}
	initDataManagers()
}

func initDataManagers() {
	if dataManager.DefineDataManager == nil {
		dataManager.DefineDataManager = &DefineDataManager{}
	}
	if dataManager.HistoricActinstDataManager == nil {
		dataManager.HistoricActinstDataManager = &HistoricActinstDataManager{}
	}
	if dataManager.HistoricIdentityLinkDataManager == nil {
		dataManager.HistoricIdentityLinkDataManager = &HistoricIdentityLinkDataManager{}
	}
	if dataManager.HistoricProcessDataManager == nil {
		dataManager.HistoricProcessDataManager = &HistoricProcessDataManager{}
	}
	if dataManager.HistoricTaskDataManager == nil {
		dataManager.HistoricTaskDataManager = &HistoricTaskDataManager{}
	}
	if dataManager.HistoricProcessDataManager == nil {
		dataManager.HistoricTaskDataManager = &HistoricTaskDataManager{}
	}
	if dataManager.HistoricVariableDataManager == nil {
		dataManager.HistoricVariableDataManager = &HistoricVariableDataManager{}
	}
	if dataManager.ProcessInstanceDataManager == nil {
		dataManager.ProcessInstanceDataManager = &ProcessInstanceDataManager{}
	}
	if dataManager.IdentityLinkDataManager == nil {
		dataManager.IdentityLinkDataManager = &IdentityLinkDataManager{}
	}
	if dataManager.TaskDataManager == nil {
		dataManager.TaskDataManager = &TaskDataManager{}
	}
	if dataManager.VariableDataManager == nil {
		dataManager.VariableDataManager = &VariableDataManager{}
	}
}

func GetDataManager() DataManager {
	return dataManager
}
