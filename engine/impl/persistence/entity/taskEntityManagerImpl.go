package entity

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity/data"
)

var (
	taskEntityImpl  TaskEntityImpl
	taskDataManager data.TaskDataManager
)

type TaskEntityManagerImpl struct {
	AbstractEntityManager
}

func (taskEntityManager TaskEntityManagerImpl) GetDataManager() data.DataManagers {
	return taskDataManager
}
func (taskEntityManager TaskEntityManagerImpl) DeleteTask(task TaskEntity) (err error) {
	taskEntityImpl.Delete(task)
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
