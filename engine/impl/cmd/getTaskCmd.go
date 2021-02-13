package cmd

import (
	"github.com/lios/go-activiti/engine/impl/interceptor"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
	"github.com/lios/go-activiti/engine/impl/persistence/entity/data"
)

type GetTaskCmd struct {
	UserId  string
	GroupId string
}

func (getTaskCmd GetTaskCmd) Execute(interceptor interceptor.CommandContext) (interface{}, error) {
	dataManager := entity.GetTaskEntityManager().GetDataManager()
	taskDataManager := dataManager.(data.TaskDataManager)
	taskResult, err := taskDataManager.QueryUndoTask(getTaskCmd.UserId, getTaskCmd.GroupId)
	if err != nil {
		return taskResult, err
	}
	return taskResult, err
}
