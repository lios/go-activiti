package entity

type AbstractEntity struct {
	AbstractEntityManager
	Id int64
}

func (entity *AbstractEntity) GetId() int64 {
	return entity.Id
}

func (entity *AbstractEntity) SetId(id int64) {
	entity.Id = id
}
