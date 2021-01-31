package persistence

type AbstractManager interface {
	Insert(data interface{}) error
}
