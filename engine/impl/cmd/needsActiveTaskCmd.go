package cmd

import (
	"github.com/lios/go-activiti/engine/impl/interceptor"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type NeedsActiveTaskCmd struct {
	AbstractTaskCmd
	TaskId       int64
	Comment      string
	TargetFlowId string
}

func (needsActiveTaskCmd NeedsActiveTaskCmd) Execute(command interceptor.CommandContext) (interface{}, error) {
	taskEntityManager := entity.GetTaskEntityManager()

	taskEntity, err := taskEntityManager.QueryTaskById(needsActiveTaskCmd.TaskId)
	if err != nil {
		return nil, err
	}
	execute, err := needsActiveTaskCmd.TaskExecute(command, taskEntity)
	return execute, err
}
