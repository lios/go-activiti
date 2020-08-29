package variable

type VariableInstanceEntit interface {
	GetType() VariableType

	SetType(variableType VariableType)
}
