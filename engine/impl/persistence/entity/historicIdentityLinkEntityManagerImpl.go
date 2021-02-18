package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	. "github.com/lios/go-activiti/model"
)

var historicIdentityLinkDataManager HistoricIdentityLinkDataManager

type HistoricIdentityLinkEntityManagerImpl struct {
	AbstractEntityManager
}

func init() {
	historicIdentityLinkDataManager = HistoricIdentityLinkDataManager{AbstractDataManager: AbstractDataManager{TableModel{AbstractModel(HistoricIdentityLink{})}}}
}

func (historicIdentityLinkEntityManagerImpl HistoricIdentityLinkEntityManagerImpl) GetDataManager() DataManagers {
	return historicIdentityLinkDataManager
}
