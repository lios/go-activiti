package engine

import (
	"github.com/lios/go-activiti/engine/behavior"
	"github.com/lios/go-activiti/engine/impl/cmd"
	. "github.com/lios/go-activiti/model"
)

type TaskServiceImpl struct {
}

//流程审批完成
func (taskService TaskServiceImpl) Complete(taskId int, variables map[string]interface{}, localScope bool) (Task, error) {
	var task Task
	exe, err := behavior.GetServiceImpl().CommandExecutor.Exe(cmd.CompleteCmd{TaskId: taskId, Variables: variables, LocalScope: localScope})
	if err != nil {
		return task, err
	}
	return exe.(Task), nil
}
