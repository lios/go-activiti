package data

import (
	"github.com/lios/go-activiti/db"
	. "github.com/lios/go-activiti/model"
	"github.com/prometheus/common/log"
	"time"
)

type HistoricActinstDataManager struct {
	AbstractDataManager
}

func (historicActinstManager HistoricActinstDataManager) FindUnfinishedHistoricActivityInstancesByExecutionAndActivityId(processInstanceId int64, actId string) (HistoricActinst, error) {
	historicActinst := HistoricActinst{}
	err := db.DB().Where("act_id = ?", actId).Where("proc_inst_id = ?", processInstanceId).First(&historicActinst).Error
	if err != nil {
		log.Infoln("Select HistoricActinst err: ", err)
		return HistoricActinst{}, err
	}
	return historicActinst, nil
}

func (historicActinstManager HistoricActinstDataManager) Update(historicActinst HistoricActinst) (err error) {
	err = db.DB().Model(&HistoricActinst{}).Where("act_id = ?", historicActinst.ActId).
		Where("proc_inst_id = ?", historicActinst.ProcessInstanceId).
		Update(&historicActinst).Error
	if err != nil {
		log.Infoln("Update HistoricActinst err: ", err)
	}
	return err
}
func (historicActinstManager HistoricActinstDataManager) UpdateProcessInstanceId(historicActinst HistoricActinst) (err error) {
	err = db.DB().Model(&HistoricActinst{}).Where("proc_inst_id = ?", historicActinst.ProcessInstanceId).
		Update(&historicActinst).Error
	if err != nil {
		log.Infoln("Update HistoricActinst err: ", err)
	}
	return err
}

func (historicActinstManager HistoricActinstDataManager) UpdateTaskId(taskId int64) (err error) {
	actinst := HistoricActinst{}
	actinst.EndTime = time.Now()
	err = db.DB().Model(&HistoricActinst{}).Where("task_id = ?", taskId).
		Update(&actinst).Error
	if err != nil {
		log.Infoln("Update HistoricActinst err: ", err)
	}
	return err
}

//func (historicActinstManager HistoricActinstDataManager) RecordTaskCreated(element model.FlowElement, entity entity2.ExecutionEntity) (err error) {
//	var actinst = HistoricActinst{}
//	actinst, err = historicActinstManager.FindUnfinishedHistoricActivityInstancesByExecutionAndActivityId(entity.GetProcessInstanceId(), element.GetId())
//	if err == nil {
//		actinst.EndTime = time.Now()
//		historicActinstManager.HistoricActinst = actinst
//		err = historicActinstManager.Update()
//	}
//	return err
//}
