package persistence

import (
	"github.com/lios/go-activiti/db"
	. "github.com/lios/go-activiti/model"
	"github.com/prometheus/common/log"
	"time"
)

type ProcessInstanceManager struct {
	Instance *ProcessInstance
}

//创建流程实例
func (processInstanceManager *ProcessInstanceManager) CreateProcessInstance() {
	err := db.DB().Create(&processInstanceManager.Instance).Error
	if err != nil {
		log.Infoln("create processInstance err", err)
	}

	processInstanceManager.createHistoricProcessInstance()
}

//查询流程实例
func (processInstanceManager *ProcessInstanceManager) GetProcessInstance(processInstanceId int64) ProcessInstance {
	instance := ProcessInstance{}
	err := db.DB().Where("id = ?", processInstanceId).Find(&instance).Error
	if err != nil {
		log.Infoln("create processInstance err", err)
	}
	return instance
}

//删除流程实例
func (processInstanceManager ProcessInstanceManager) DeleteProcessInstance(processInstanceId int64) (err error) {
	err = db.DB().Where("id = ?", processInstanceId).Delete(&ProcessInstance{}).Error
	if err != nil {
		log.Infoln("delete processInstance err ", err)
		return err
	}
	err = processInstanceManager.recordActivityEnd(processInstanceId)
	return err
}

func (processInstanceManager ProcessInstanceManager) recordActivityEnd(processInstanceId int64) (err error) {
	historicProcessManager := HistoricProcessManager{}
	historicProcess := HistoricProcess{}
	historicProcess.ProcessInstanceId = processInstanceId
	historicProcess.EndTime = time.Now()
	historicProcessManager.HistoricProcess = historicProcess
	err = historicProcessManager.MarkEnded()
	return err
}

func (processInstanceManager *ProcessInstanceManager) createHistoricProcessInstance() (err error) {
	processInstance := processInstanceManager.Instance
	historicProcess := HistoricProcess{}
	//historicProcess.ProcessInstanceEntity = processInstance.ProcessInstanceEntity
	historicProcess.ProcessInstanceId = processInstance.Id
	historicProcess.DeploymentId = processInstance.DeploymentId
	historicProcess.TenantId = processInstance.TenantId
	historicProcess.StartTime = processInstance.StartTime
	historicProcess.Name = processInstance.Name
	historicProcess.BusinessKey = processInstance.BusinessKey
	historicProcess.StartUserId = processInstance.StartUserId
	historicProcess.Key = processInstance.Key
	historicProcess.ProcessDefineId = processInstance.ProcessDefineId

	historicProcessManager := HistoricProcessManager{}
	historicProcessManager.HistoricProcess = historicProcess
	return historicProcessManager.Insert()
}
