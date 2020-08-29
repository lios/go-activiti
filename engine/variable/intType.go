package variable

type IntType struct {
}

func (intType IntType) GetTypeName() string {
	return "int"
}

func (intType IntType) GetValue(valueFields ValueFields) interface{} {
	return valueFields.GetNumberValue()
}

func (intType IntType) SetValue(value interface{}, valueFields ValueFields) {
	valueInt, ok := value.(int)
	if ok {
		valueFields.SetNumberValue(valueInt)
	}
	valueFields.SetBlobValue("")
}
