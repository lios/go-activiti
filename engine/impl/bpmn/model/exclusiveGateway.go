package model

type ExclusiveGateway struct {
	Gateway
}

func (exclusiveGateway ExclusiveGateway) GetType() string {
	return "ExclusiveGateway"
}
