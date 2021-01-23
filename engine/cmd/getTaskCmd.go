package cmd

import (
	"github.com/lios/go-activiti/engine/behavior"
	. "github.com/lios/go-activiti/engine/manager"
)

type GetTaskCmd struct {
	UserId  string
	GroupId string
}

func (getTaskCmd GetTaskCmd) Execute(interceptor behavior.CommandContext) (interface{}, error) {
	manager := GetTaskManager()
	taskResult, err := manager.QueryUndoTask(getTaskCmd.UserId, getTaskCmd.GroupId)
	if err != nil {
		return taskResult, err
	}
	return taskResult, err
}
