package variable

type StringType struct {
}

func (stringType StringType) GetTypeName() string {
	return "string"
}

func (stringType StringType) GetValue(valueFields ValueFields) interface{} {
	return valueFields.GetTextValue()
}

func (stringType StringType) SetValue(value interface{}, valueFields ValueFields) {
	b, ok := value.(string)
	if ok {
		valueFields.SetTextValue(b)
	}
	valueFields.SetBlobValue("")
}
