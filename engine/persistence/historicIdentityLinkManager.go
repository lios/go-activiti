package persistence

import (
	"github.com/lios/go-activiti/db"
	"github.com/lios/go-activiti/model"
	"github.com/prometheus/common/log"
)

type HistoricIdentityLinkManager struct {
	HistoricIdentityLink model.HistoricIdentityLink
}

func (historicIdentityLink HistoricIdentityLinkManager) Insert() (err error) {
	err = db.DB().Create(&historicIdentityLink.HistoricIdentityLink).Error
	if err != nil {
		log.Infoln("Create HistoricIdentityLink Err", err)
	}
	return err
}
