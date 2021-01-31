package data

import (
	"github.com/lios/go-activiti/db"
	. "github.com/lios/go-activiti/model"
	"github.com/prometheus/common/log"
)

type HistoricProcessDataManager struct {
	AbstractDataManager
	HistoricProcess HistoricProcess
}

func (historicProcessManager HistoricProcessDataManager) Insert() (err error) {
	err = db.DB().Create(&historicProcessManager.HistoricProcess).Error
	if err != nil {
		log.Infoln("Create HistoricActinst Err", err)
	}
	return err
}

func (historicProcessManager HistoricProcessDataManager) MarkEnded() (err error) {
	historicProcess := historicProcessManager.HistoricProcess
	err = db.DB().Where("proc_inst_id=?", historicProcess.ProcessInstanceId).Update(&historicProcess).Error
	if err != nil {
		log.Infoln("delete HistoricProcess Err", err)
		return err
	}
	historicActinst := HistoricActinst{}
	historicActinst.EndTime = historicProcess.EndTime
	historicProcess.ProcessInstanceId = historicProcess.Id
	historicActinstManager := HistoricActinstDataManager{}
	historicActinstManager.HistoricActinst = historicActinst
	err = historicActinstManager.UpdateProcessInstanceId()
	return err
}
