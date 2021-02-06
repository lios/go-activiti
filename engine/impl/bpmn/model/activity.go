package model

type Activity struct {
	*FlowNode
	LoopCharacteristics *MultiInstanceLoopCharacteristics
}

func (activity Activity) GetLoopCharacteristics() *MultiInstanceLoopCharacteristics {
	return activity.LoopCharacteristics
}

func (activity Activity) SetLoopCharacteristics(loopCharacteristics *MultiInstanceLoopCharacteristics) {
	activity.LoopCharacteristics = loopCharacteristics
}
