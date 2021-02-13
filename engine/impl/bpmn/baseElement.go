package bpmn

//通用字段
type BaseElement interface {
	GetId() string
	GetName() string
}

type BaseHandlerType interface {
	GetType() string
}
