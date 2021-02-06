package persistence

type AbstractManager interface {
	FindById(entityId string)
	Insert(data interface{}) error
}
