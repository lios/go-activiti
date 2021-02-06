package parse

import . "github.com/lios/go-activiti/engine/impl/persistence/entity"

type ParsedDeployment struct {
	BpmnParse                        BpmnParse
	ProcessDefinitionEntity          ProcessDefinitionEntity
	MapProcessDefinitionsToParses    map[ProcessDefinitionEntity]BpmnParse
	MapProcessDefinitionsToResources map[ProcessDefinitionEntity]ResourceEntity
}
