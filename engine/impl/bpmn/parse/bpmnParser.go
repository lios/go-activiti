package parse

import (
	factory2 "github.com/lios/go-activiti/engine/impl/bpmn/parse/factory"
)

type BpmnParser struct {
	ActivityBehaviorFactory factory2.ActivityBehaviorFactory
	BpmnParserHandlers      BpmnParseHandlers
	BpmnParseFactory        BpmnParseFactory
}

func (bpmnParser BpmnParser) CreateParse() BpmnParse {
	return bpmnParser.BpmnParseFactory.CreateBpmnParse(bpmnParser)
}

func (bpmnParser BpmnParser) SetActivityBehaviorFactory(activityBehaviorFactory factory2.ActivityBehaviorFactory) {
	bpmnParser.ActivityBehaviorFactory = activityBehaviorFactory
}

func (bpmnParser BpmnParser) GetActivityBehaviorFactory() factory2.ActivityBehaviorFactory {
	return bpmnParser.ActivityBehaviorFactory
}

func (bpmnParser BpmnParser) GetBpmnParserHandlers() BpmnParseHandlers {
	return bpmnParser.BpmnParserHandlers
}

func (bpmnParser BpmnParser) SetBpmnParserHandlers(bpmnParseHandlers BpmnParseHandlers) {
	bpmnParser.BpmnParserHandlers = bpmnParseHandlers
}
