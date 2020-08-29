package variable

type BooleanType struct {
}

func (boolType BooleanType) GetTypeName() string {
	return "bool"
}

func (boolType BooleanType) GetValue(valueFields ValueFields) interface{} {
	return valueFields.GetTextValue() == "true"
}

func (boolType BooleanType) SetValue(value interface{}, valueFields ValueFields) {
	b, ok := value.(bool)
	if ok {
		if b {
			valueFields.SetTextValue("true")
		} else {
			valueFields.SetTextValue("false")
		}
	}
	valueFields.SetBlobValue("")
}
