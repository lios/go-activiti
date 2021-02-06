package data

import (
	"github.com/lios/go-activiti/db"
	"github.com/lios/go-activiti/logger"
	. "github.com/lios/go-activiti/model"
	"github.com/prometheus/common/log"
	"time"
)

type HistoricTaskDataManager struct {
	HistoricTask
	AbstractDataManager
	HistoricTaskModel HistoricTask
}

func (historicTaskManager HistoricTaskDataManager) GetByTaskId(id int64) (HistoricTask, error) {
	task := HistoricTask{}
	err := db.DB().Where("task_id = ?", id).First(&task).Error
	if err != nil {
		logger.Error("find bu id err:", err)
	}
	return task, err
}

func (historicTaskManager HistoricTaskDataManager) MarkEnded(id int64) (err error) {
	historicTask := HistoricTask{}
	historicTask.EndTime = time.Now()
	err = db.DB().Model(&HistoricTask{}).Where("id=?", id).Update(&historicTask).Error
	if err != nil {
		log.Infoln("Update HistoricTask Err", err)
	}
	return err
}
