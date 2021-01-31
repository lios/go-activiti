package parse

type BpmnParseFactory interface {
	CreateBpmnParse(bpmnParser BpmnParser) BpmnParse
}
