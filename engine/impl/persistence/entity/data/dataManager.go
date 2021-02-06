package data

type DataManagers interface {
	Insert(data interface{}) error

	FindById(id int64, data interface{}) error

	Delete(id int64) error

	GetTableName() string
}
