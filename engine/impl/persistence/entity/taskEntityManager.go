package entity

type TaskEntityManager interface {
	EntityManager

	QueryTaskById(id int64) (task TaskEntity, err error)

	DeleteTask(task TaskEntity) (err error)

	FindByProcessInstanceId(processInstanceId int64) (taskEntity []TaskEntityImpl, err error)
}
