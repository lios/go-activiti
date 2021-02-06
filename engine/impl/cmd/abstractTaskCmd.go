package cmd

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
	"github.com/lios/go-activiti/engine/interceptor"
)

type AbstractTaskCmd interface {
	TaskExecute(command interceptor.CommandContext, entity entity.TaskEntity) (interface{}, error)
}
