package entity

type HistoricTaskInstanceEntityManager interface {
	EntityManager
	delete(taskId string)
}
