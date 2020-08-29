package variable

import "encoding/json"

type MapType struct {
}

func (mapType MapType) GetTypeName() string {
	return "map"
}

func (mapType MapType) GetValue(valueFields ValueFields) interface{} {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(valueFields.GetTextValue()), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}

func (mapType MapType) SetValue(value interface{}, valueFields ValueFields) {
	b, ok := value.(map[string]interface{})
	if ok {
		dataType, _ := json.Marshal(b)
		dataString := string(dataType)
		valueFields.SetTextValue(dataString)
	}
}
