package cfg

import (
	. "github.com/lios/go-activiti/engine/impl/cmd"
	"github.com/lios/go-activiti/engine/task"
	. "github.com/lios/go-activiti/model"
)

type TaskServiceImpl struct {
	ServiceImpl
}

//查询待审批任务
func (taskService TaskServiceImpl) QueryUndoTask(userId, groupId string) ([]task.TaskInfo, error) {
	exe, err := taskService.GetCommandExecutor().Exe(GetTaskCmd{UserId: userId, GroupId: groupId})
	if err != nil {
		return nil, err
	}
	results, ok := exe.([]task.TaskInfo)
	if ok {
		return results, nil
	}
	return nil, err
}

//流程审批完成
func (taskService TaskServiceImpl) Complete(taskId int64, variables map[string]interface{}, localScope bool) (Task, error) {
	var task Task
	exe, err := taskService.GetCommandExecutor().Exe(CompleteCmd{NeedsActiveTaskCmd: NeedsActiveTaskCmd{AbstractTaskCmd: AbstractTaskCmd(CompleteCmd{Variables: variables, LocalScope: localScope}), TaskId: taskId}})
	if err != nil {
		return task, err
	}
	return exe.(Task), nil
}

//查询待审批任务
func (taskService TaskServiceImpl) BackTask(taskId int64, targetFlowId string) (bool, error) {
	exe, err := taskService.GetCommandExecutor().Exe(BackTaskCmd{NeedsActiveTaskCmd: NeedsActiveTaskCmd{AbstractTaskCmd: AbstractTaskCmd(BackTaskCmd{TargetFlowId: targetFlowId}), TaskId: taskId}})
	if err != nil {
		return false, err
	}
	return exe.(bool), nil
}
