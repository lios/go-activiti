package entity

import . "github.com/lios/go-activiti/engine/impl/persistence/entity/data"

type AbstractEntityManager interface {
	GetTaskEntityManager() TaskDataManager

	GetDefineEntityManager() DefineDataManager

	GetVariableEntityManager() VariableDataManager

	GetIdentityLinkEntityManager() IdentityLinkDataManager

	GetHistoricActinstEntityManager() HistoricActinstDataManager

	GetHistoricTaskEntityManager() HistoricTaskDataManager

	GetHistoricProcessEntityManager() HistoricProcessDataManager
}
