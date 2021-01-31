package model

type MultiInstanceLoopCharacteristics struct {
	IsSequential        bool   `xml:"isSequential,attr"`
	Collection          string `xml:"collection,attr"`
	CompletionCondition string `xml:"completionCondition"`
}
