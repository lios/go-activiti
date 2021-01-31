package engine

import (
	"github.com/lios/go-activiti/engine/task"
	"github.com/lios/go-activiti/model"
)

type TaskService interface {
	QueryUndoTask(userId, groupId string) ([]task.TaskInfo, error)

	Complete(taskId int, variables map[string]interface{}, localScope bool) (model.Task, error)

	BackTask(taskId int, targetFlowId string) (bool, error)
}
