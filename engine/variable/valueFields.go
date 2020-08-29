package variable

type ValueFields interface {
	GetName() string

	GetProcessInstanceId() int64

	GetTaskId() int64

	GetNumberValue() int

	SetNumberValue(value int)

	GetTextValue() string

	SetTextValue(value string)

	SetBlobValue(value string)

	GetBlobValue() string
}
