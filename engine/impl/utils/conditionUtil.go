package utils

import (
	. "github.com/heartlhj/go-expression/expression"
	. "github.com/heartlhj/go-expression/expression/spel"
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

var (
	context = StandardEvaluationContext{}
	parser  = SpelExpressionParser{}
)

type ConditionUtil struct {
}

func HasTrueCondition(sequenceFlow model.SequenceFlow, execution entity.ExecutionEntity) bool {
	var conditionExpression = sequenceFlow.ConditionExpression
	if conditionExpression != "" {
		variable := execution.GetProcessVariable()
		context.SetVariables(variable)
		valueContext := parser.ParseExpression(conditionExpression).GetValueContext(&context)
		b, ok := valueContext.(bool)
		if ok {
			return b
		}
		return false
	} else {
		return true
	}

}
