package persistence

import (
	"github.com/lios/go-activiti/db"
	. "github.com/lios/go-activiti/model"
	"github.com/prometheus/common/log"
)

type HistoricProcessManager struct {
	HistoricProcess HistoricProcess
}

func (historicProcessManager HistoricProcessManager) Insert() (err error) {
	err = db.DB().Create(&historicProcessManager.HistoricProcess).Error
	if err != nil {
		log.Infoln("Create HistoricActinst Err", err)
	}
	return err
}

func (historicProcessManager HistoricProcessManager) MarkEnded() (err error) {
	historicProcess := historicProcessManager.HistoricProcess
	err = db.DB().Where("proc_inst_id=?", historicProcess.ProcessInstanceId).Update(&historicProcess).Error
	if err != nil {
		log.Infoln("delete HistoricProcess Err", err)
		return err
	}
	historicActinst := HistoricActinst{}
	historicActinst.EndTime = historicProcess.EndTime
	historicProcess.ProcessInstanceId = historicProcess.Id
	historicActinstManager := HistoricActinstManager{}
	historicActinstManager.HistoricActinst = historicActinst
	err = historicActinstManager.UpdateProcessInstanceId()
	return err
}
