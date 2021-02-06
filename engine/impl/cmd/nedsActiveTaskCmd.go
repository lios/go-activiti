package cmd

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
	"github.com/lios/go-activiti/engine/interceptor"
)

type NeedsActiveTaskCmd struct {
	AbstractTaskCmd
	TaskId int64
}

func (needsActiveTaskCmd NeedsActiveTaskCmd) Execute(command interceptor.CommandContext) (interface{}, error) {
	taskEntityManager := entity.GetTaskEntityManager()
	task := taskEntityManager.GetById(needsActiveTaskCmd.TaskId)
	taskEntity := task.(entity.TaskEntity)
	execute, err := needsActiveTaskCmd.TaskExecute(command, taskEntity)
	return execute, err
}
