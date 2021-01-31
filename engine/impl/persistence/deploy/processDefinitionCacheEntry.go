package deploy

import (
	"github.com/lios/go-activiti/engine/impl/bpmn/model"
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

type ProcessDefinitionCacheEntry struct {
	Process           model.Process
	ProcessDefinition entity.ProcessInstanceEntity
}

func NewProcessDefinitionCacheEntry(process model.Process) ProcessDefinitionCacheEntry {
	return ProcessDefinitionCacheEntry{Process: process}
}
