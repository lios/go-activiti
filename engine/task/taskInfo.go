package task

import "time"

type TaskInfo struct {
	Id             int64     `json:"id"`
	TaskDefineKey  string    `json:"task_define_key"`
	TaskDefineName string    `json:"task_define_name"`
	TenantId       string    `json:"tenant_id"`
	DeploymentId   int       `json:"deployment_id"`
	StartTime      time.Time `json:"start_time"`
	Assignee       string    `json:"assignee"`
	ProcInstId     int64     `json:"proc_inst_id"`
	UserId         string    `json:"user_id"`
	GroupId        string    `json:"group_id"`
}
