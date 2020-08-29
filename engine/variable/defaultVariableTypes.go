package variable

var (
	typesList = make(map[string]VariableType, 0)
)

type DefaultVariableTypes struct {
}

func (variableTypes DefaultVariableTypes) AddType(variableType VariableType) {
	typesList[variableType.GetTypeName()] = variableType
}

func (variableTypes DefaultVariableTypes) GetVariableType(typeName string) VariableType {
	return typesList[typeName]
}
