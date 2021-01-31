package interceptor

import (
	"container/list"
	"github.com/lios/go-activiti/errs"
)

var (
	CommandStack *list.List
)

func init() {
	CommandStack = list.New()
}

type Context struct {
}

func SetCommandContext(commandContext CommandContext) {
	CommandStack.PushFront(commandContext)
}

func GetCommandContext() (CommandContext, error) {
	if CommandStack.Len() <= 0 {
		return CommandContext{}, errs.ProcessError{}
	}
	return CommandStack.Front().Value.(CommandContext), nil
}

func GetAgenda() CommandContext {
	commandContext, _ := GetCommandContext()
	return commandContext
}
