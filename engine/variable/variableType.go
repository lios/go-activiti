package variable

type VariableType interface {
	GetTypeName() string

	GetValue(valueFields ValueFields) interface{}

	SetValue(value interface{}, valueFields ValueFields)
}
