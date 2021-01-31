package parse

type DefaultBpmnParseFactory struct {
}

func CreateBpmnParse(bpmnParser BpmnParser) BpmnParse {
	return BpmnParse{ActivityBehaviorFactory: bpmnParser.ActivityBehaviorFactory,
		BpmnParserHandlers: bpmnParser.BpmnParserHandlers}
}
