package entity

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity/data"
)

var (
	taskEntityImpl  TaskEntityImpl
	taskDataManager TaskDataManager
)

type TaskEntityManagerImpl struct {
	AbstractEntityManager
}

func (taskEntityManager TaskEntityManagerImpl) GetDataManager() DataManagers {
	return taskDataManager
}
func (taskEntityManager TaskEntityManagerImpl) DeleteTask(task TaskEntity) (err error) {
	taskEntityManager.Delete(task)
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
