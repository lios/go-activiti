package entity

type AbstractEntityManager struct {
	EntityManager
}

func (entityManager AbstractEntityManager) Insert(data interface{}) error {
	return entityManager.GetDataManager().Insert(data)
}

func (entityManager AbstractEntityManager) GetById(id int64, data interface{}) interface{} {
	entityManager.GetDataManager().FindById(id, data)
	return entityManager.GetDataManager().FindById(id, data)
}

func (entityManager AbstractEntityManager) Delete(entity Entity) {
	entityManager.GetDataManager().Delete(entity.GetId())
}
