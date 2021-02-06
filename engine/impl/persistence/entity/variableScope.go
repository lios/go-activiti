package entity

type VariableScope interface {
	SetVariableLocal(parameters map[string]interface{}) error
}
