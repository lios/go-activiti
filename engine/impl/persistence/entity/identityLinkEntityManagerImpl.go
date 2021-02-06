package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	"github.com/lios/go-activiti/logger"
)

var (
	identityLinkEntity      IdentityLinkEntity
	identityLinkDataManager IdentityLinkDataManager
)

type IdentityLinkEntityManagerImpl struct {
	AbstractEntityManager
}

func (identityLinkEntityManager IdentityLinkEntityManagerImpl) GetDataManager() DataManagers {
	return identityLinkDataManager
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
