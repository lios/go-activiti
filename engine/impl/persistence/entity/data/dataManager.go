package data

type DataManager interface {
	Insert(data interface{}) error
}
