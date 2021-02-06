package variable

var variableManager VariableManager

type VariableManager struct {
	VariableTypes VariableTypes
}

func init() {
	variableManager = VariableManager{}
	initVariableTypes()
}

func initVariableTypes() {
	defaultVariableTypes := DefaultVariableTypes{}
	defaultVariableTypes.AddType(BooleanType{})
	defaultVariableTypes.AddType(IntType{})
	defaultVariableTypes.AddType(StringType{})
	defaultVariableTypes.AddType(MapType{})
	variableManager.VariableTypes = defaultVariableTypes
}

func GetVariableManager() VariableManager {
	return variableManager
}
