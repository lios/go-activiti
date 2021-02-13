package cmd

import (
	"github.com/lios/go-activiti/engine/impl/interceptor"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type AbstractTaskCmd interface {
	TaskExecute(command interceptor.CommandContext, entity entity.TaskEntity) (interface{}, error)
}
