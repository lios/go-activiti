package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	"github.com/lios/go-activiti/logger"
	. "github.com/lios/go-activiti/model"
)

var (
	taskEntityImpl  TaskEntityImpl
	taskDataManager TaskDataManager
)

func init() {
	taskDataManager = TaskDataManager{AbstractDataManager: AbstractDataManager{TableModel{AbstractModel(Task{})}}}
}

type TaskEntityManagerImpl struct {
	AbstractEntityManager
}

func (taskEntityManager TaskEntityManagerImpl) GetDataManager() DataManagers {
	return taskDataManager
}
func (taskEntityManager TaskEntityManagerImpl) QueryTaskById(id int64) (taskEntity TaskEntity, err error) {
	manager := taskEntityManager.GetDataManager()
	task := Task{}
	err = manager.FindById(id, &task)
	if err != nil {
		logger.Error("query task err :", err)
		return nil, err
	}
	entityImpl := TaskEntityImpl{}
	entityImpl.SetStartTime(task.StartTime)
	entityImpl.SetProcessInstanceId(task.ProcessInstanceId)
	entityImpl.SetId(task.Id)
	entityImpl.SetTaskDefineKey(task.TaskDefineKey)
	entityImpl.SetTaskDefineName(task.TaskDefineName)
	return &entityImpl, nil
}
func (taskEntityManager TaskEntityManagerImpl) DeleteTask(task TaskEntity) (err error) {
	manager := taskEntityManager.GetDataManager()
	dataManager := manager.(TaskDataManager)
	dataManager.Delete(task.GetId())
	if err != nil {
		return err
	}
	identityLinkManager := GetIdentityLinkManager()
	identityLinkManager.DeleteIdentityLinksByTaskId(task.GetId())

	variableManager := GetVariableEntityManager()
	variableManager.DeleteVariableInstanceByTask(task.GetId())
	err = taskEntityManager.recordTaskEnd(task)
	return err
}

func (taskEntityManager TaskEntityManagerImpl) recordTaskEnd(task TaskEntity) (err error) {
	historicTaskManager := GetHistoricTaskInstanceEntityManager()
	historicTaskManager.RecordTaskEnd(task.GetId(), "")

	activityInstanceEntityManager := GetHistoricActivityInstanceEntityManager()
	activityInstanceEntityManager.RecordEnd(task.GetId())

	return nil
}

func (taskEntityManager TaskEntityManagerImpl) FindByProcessInstanceId(processInstanceId int64) (taskEntity []TaskEntityImpl, err error) {
	manager := taskEntityManager.GetDataManager()
	dataManager := manager.(TaskDataManager)
	tasks, err := dataManager.FindByProcessInstanceId(processInstanceId)
	if err != nil {
		return taskEntity, err
	}
	taskEntitys := make([]TaskEntityImpl, 0)
	for _, task := range tasks {
		taskEntity := TaskEntityImpl{}
		taskEntity.SetId(task.Id)
		taskEntitys = append(taskEntitys, taskEntity)
	}
	return taskEntitys, nil
}

func (taskEntityManager TaskEntityManagerImpl) InsertTask(taskEntity *TaskEntityImpl) error {
	//dataManager := taskEntityManager.GetDataManager()
	task := Task{}
	task.ProcessInstanceId = taskEntity.GetProcessInstanceId()
	task.TaskDefineKey = taskEntity.taskDefineKey
	task.TaskDefineName = taskEntity.taskDefineName
	task.StartTime = taskEntity.startTime
	err := taskEntityManager.Insert(&task)
	taskEntity.SetId(task.Id)
	if err != nil {
		logger.Error("create task err:", err)
		return err
	}
	taskEntityManager.recordTaskCreated(task)
	return err
}

func (taskEntityManager TaskEntityManagerImpl) recordTaskCreated(task Task) (err error) {
	manager := GetHistoricTaskInstanceEntityManager().GetDataManager()
	historicTaskManager := manager.(HistoricTaskDataManager)
	historicTask := taskEntityManager.createHistoricTask(task)
	err = historicTaskManager.Insert(&historicTask)
	if err != nil {
		logger.Error("create task err:", err)
		return err
	}
	entityManager := GetHistoricActivityInstanceEntityManager()
	entityManager.RecordTaskId(task)
	return err
}

func (taskEntityManager TaskEntityManagerImpl) createHistoricTask(task Task) HistoricTask {
	historicTask := HistoricTask{}
	//historicTask.TaskEntity = task.TaskEntity
	historicTask.TaskId = task.Id
	historicTask.ProcessInstanceId = task.ProcessInstanceId
	historicTask.StartTime = task.StartTime
	historicTask.TenantId = task.TenantId
	historicTask.Assignee = task.Assignee
	historicTask.TaskDefineKey = task.TaskDefineKey
	historicTask.DeploymentId = task.DeploymentId
	historicTask.TaskDefineName = task.TaskDefineName
	return historicTask
}
