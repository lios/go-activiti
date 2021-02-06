package entity

type VariableEntityImpl struct {
	AbstractEntity
	Type   string
	Delete bool
}

func (variableEntity VariableEntityImpl) GetType() string {
	return variableEntity.Type
}

func (variableEntity VariableEntityImpl) SetDeleted(b bool) {
	variableEntity.Delete = b
}
