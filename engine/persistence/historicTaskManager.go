package persistence

import (
	"github.com/lios/go-activiti/db"
	. "github.com/lios/go-activiti/model"
	"github.com/prometheus/common/log"
)

type HistoricTaskManager struct {
	HistoricTask HistoricTask
}

func (historicTaskManager HistoricTaskManager) Insert() (err error) {
	err = db.DB().Create(&historicTaskManager.HistoricTask).Error
	if err != nil {
		log.Infoln("Create HistoricTask Err", err)
	}
	return err
}

func (historicTaskManager HistoricTaskManager) MarkEnded() (err error) {
	err = db.DB().Model(&HistoricTask{}).Where("task_id=?", historicTaskManager.HistoricTask.TaskId).Update(&historicTaskManager.HistoricTask).Error
	if err != nil {
		log.Infoln("Update HistoricTask Err", err)
	}
	return err
}
