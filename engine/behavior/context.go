package behavior

import (
	"container/list"
	"github.com/lios/go-activiti/engine"
	"github.com/lios/go-activiti/errs"
)

var (
	Stack list.List
)

type Context struct {
}

func SetCommandContext(commandContext CommandContext) {
	Stack.PushFront(commandContext)
}

func GetCommandContext() (CommandContext, error) {
	if Stack.Len() <= 0 {
		return CommandContext{}, errs.ProcessError{}
	}
	return Stack.Front().Value.(CommandContext), nil
}

func GetAgenda() engine.ActivitiEngineAgenda {
	return Stack.Front().Value.(CommandContext).Agenda
}
