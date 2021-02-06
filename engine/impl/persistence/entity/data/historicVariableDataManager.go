package data

import (
	"github.com/lios/go-activiti/db"
	"github.com/lios/go-activiti/engine/variable"
	"github.com/prometheus/common/log"
)

type HistoricVariableDataManager struct {
	DataManagers
	HistoricVariable variable.HistoricVariable
}

func (historicVariableManager HistoricVariableDataManager) Insert() (err error) {
	err = db.DB().Create(&historicVariableManager.HistoricVariable).Error
	if err != nil {
		log.Infoln("Create HistoricVariable Err ", err)
	}
	return err
}
