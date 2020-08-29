package engine

import . "github.com/lios/go-activiti/model"

type TaskService interface {
	Complete(taskId int, variables map[string]interface{}) Task
}
