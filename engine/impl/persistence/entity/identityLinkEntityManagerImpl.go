package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	"github.com/lios/go-activiti/logger"
	. "github.com/lios/go-activiti/model"
	"github.com/prometheus/common/log"
)

var (
	identityLinkEntity      IdentityLinkEntity
	identityLinkDataManager IdentityLinkDataManager
)

func init() {
	identityLinkDataManager = IdentityLinkDataManager{AbstractDataManager: AbstractDataManager{TableModel{AbstractModel(IdentityLink{})}}}
}

type IdentityLinkEntityManagerImpl struct {
	AbstractEntityManager
}

func (identityLinkEntityManager IdentityLinkEntityManagerImpl) GetDataManager() DataManagers {
	return identityLinkDataManager
}

func (identityLinkEntityManager IdentityLinkEntityManagerImpl) CreateIdentityLink(identityLink IdentityLink) (err error) {
	manager := identityLinkEntityManager.GetDataManager()
	linkDataManager := manager.(IdentityLinkDataManager)
	linkDataManager.Insert(&identityLink)
	if err != nil {
		log.Infoln("Create IdentityLink Err ", err)
		return err
	}
	err = identityLinkManager.createHistoricIdentityLink(identityLink)
	return err
}

func (identityLinkEntityManager IdentityLinkEntityManagerImpl) createHistoricIdentityLink(identityLink IdentityLink) (err error) {

	historicIdentityLink := HistoricIdentityLink{}
	historicIdentityLink.UserId = identityLink.UserId
	historicIdentityLink.TaskId = identityLink.TaskId
	historicIdentityLink.ProcessInstanceId = identityLink.ProcessInstanceId
	manager := GetHistoricIdentityLinkEntityManager().GetDataManager()
	linkDataManager := manager.(HistoricIdentityLinkDataManager)
	err = linkDataManager.Insert(&historicIdentityLink)
	return err
}

func (identityLinkEntityManager IdentityLinkEntityManagerImpl) DeleteIdentityLinksByTaskId(taskId int64) {
	dataManager := identityLinkEntityManager.GetDataManager()
	linkDataManager := dataManager.(IdentityLinkDataManager)
	links, err := linkDataManager.SelectByTaskId(taskId)
	if err != nil {
		logger.Error("select by taskId err:", err)
		return
	}
	for _, link := range links {
		linkDataManager.Delete(link.Id)
	}
}
