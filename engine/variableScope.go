package engine

type VariableScope interface {
	SetVariable(variableName string, value interface{})
}
