package parse

import "github.com/lios/go-activiti/engine/impl/persistence/entity"

type DefaultBpmnParseFactory struct {
}

func (DefaultBpmnParseFactory) CreateBpmnParse(bpmnParser BpmnParser) BpmnParse {
	return BpmnParse{ActivityBehaviorFactory: bpmnParser.ActivityBehaviorFactory,
		BpmnParserHandlers: bpmnParser.BpmnParserHandlers, ProcessDefinitions: make([]entity.ProcessDefinitionEntity, 0)}
}
