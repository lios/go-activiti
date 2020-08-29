package variable

type VariableTypes interface {
	AddType(variableType VariableType)

	GetVariableType(typeName string) VariableType
}
