package entity

type VariableEntity interface {
	GetType() string

	SetDeleted(b bool)
}
