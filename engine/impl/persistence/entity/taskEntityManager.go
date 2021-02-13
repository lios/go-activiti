package entity

type TaskEntityManager interface {
	EntityManager
	DeleteTask(task TaskEntity) (err error)

	FindByProcessInstanceId(processInstanceId int64) (taskEntity []TaskEntityImpl, err error)
}
