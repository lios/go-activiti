package cmd

import (
	"github.com/lios/go-activiti/engine/interceptor"
)

type GetTaskCmd struct {
	UserId  string
	GroupId string
}

func (getTaskCmd GetTaskCmd) Execute(interceptor interceptor.CommandContext) (interface{}, error) {
	manager := interceptor.ProcessEngineConfiguration.TaskDataManager
	taskResult, err := manager.QueryUndoTask(getTaskCmd.UserId, getTaskCmd.GroupId)
	if err != nil {
		return taskResult, err
	}
	return taskResult, err
}
