package engine

import (
	"github.com/lios/go-activiti/engine/behavior"
	"github.com/lios/go-activiti/engine/cmd"
	. "github.com/lios/go-activiti/entity"
	. "github.com/lios/go-activiti/model"
)

type TaskService struct {
}

//查询待审批任务
func (taskService TaskService) QueryUndoTask(userId, groupId string) ([]TaskEntity, error) {
	exe, err := behavior.GetServiceImpl().CommandExecutor.Exe(cmd.GetTaskCmd{UserId: userId, GroupId: groupId})
	if err != nil {
		return nil, err
	}
	results, ok := exe.([]TaskEntity)
	if ok {
		return results, nil
	}
	return nil, err
}

//流程审批完成
func (taskService TaskService) Complete(taskId int, variables map[string]interface{}, localScope bool) (Task, error) {
	var task Task
	exe, err := behavior.GetServiceImpl().CommandExecutor.Exe(cmd.CompleteCmd{TaskId: taskId, Variables: variables, LocalScope: localScope})
	if err != nil {
		return task, err
	}
	return exe.(Task), nil
}

//查询待审批任务
func (taskService TaskService) BackTask(taskId int, targetFlowId string) (bool, error) {
	exe, err := behavior.GetServiceImpl().CommandExecutor.Exe(cmd.BackTaskCmd{TaskId: taskId, TargetFlowId: targetFlowId})
	if err != nil {
		return false, err
	}
	return exe.(bool), nil
}
