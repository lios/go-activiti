package engine

import (
	. "github.com/lios/go-activiti/engine/impl/persistence/entity"
	"github.com/lios/go-activiti/engine/task"
)

type TaskService interface {
	QueryUndoTask(userId, groupId string) ([]task.TaskInfo, error)

	Complete(taskId int64, variables map[string]interface{}, localScope bool) (TaskEntity, error)

	BackTask(taskId int64, targetFlowId string) (bool, error)
}
