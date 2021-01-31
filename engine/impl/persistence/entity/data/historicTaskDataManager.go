package data

import (
	"github.com/lios/go-activiti/db"
	. "github.com/lios/go-activiti/model"
	"github.com/prometheus/common/log"
)

type HistoricTaskDataManager struct {
	HistoricTask HistoricTask
}

func (historicTaskManager HistoricTaskDataManager) Insert() (err error) {
	err = db.DB().Create(&historicTaskManager.HistoricTask).Error
	if err != nil {
		log.Infoln("Create HistoricTask Err", err)
	}
	return err
}

func (historicTaskManager HistoricTaskDataManager) MarkEnded() (err error) {
	err = db.DB().Model(&HistoricTask{}).Where("task_id=?", historicTaskManager.HistoricTask.TaskId).Update(&historicTaskManager.HistoricTask).Error
	if err != nil {
		log.Infoln("Update HistoricTask Err", err)
	}
	return err
}
