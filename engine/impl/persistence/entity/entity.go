package entity

type Entity interface {
	GetId() int64

	SetId(id int64)
}
