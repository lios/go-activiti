package entity

type TaskEntityManager interface {
	EntityManager
	DeleteTask(task TaskEntity) (err error)
}
