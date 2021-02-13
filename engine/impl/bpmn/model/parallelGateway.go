package model

type ParallelGateway struct {
	Gateway
}

func (parallelGateway ParallelGateway) GetType() string {
	return "ParallelGateway"
}
