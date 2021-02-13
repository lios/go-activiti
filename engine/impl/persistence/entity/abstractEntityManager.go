package entity

type AbstractEntityManager struct {
	EntityManager
}

func (entityManager AbstractEntityManager) Insert(data interface{}) error {
	return entityManager.GetDataManager().Insert(data)
}

func (entityManager AbstractEntityManager) GetById(id int64) Entity {
	return &AbstractEntity{}
}

func (entityManager AbstractEntityManager) Delete(entity Entity) {
	entityManager.GetDataManager().Delete(entity.GetId())
}
