package model

type InclusiveGateway struct {
	Gateway
}

func (inclusiveGateway InclusiveGateway) GetType() string {
	return "InclusiveGateway"
}
