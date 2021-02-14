package deploy

import "sync"

var deploymentCache *sync.Map

type DeploymentCache interface {
	Get(id string) ProcessDefinitionCacheEntry
	Add(id string, object ProcessDefinitionCacheEntry)
	Clear()
}
