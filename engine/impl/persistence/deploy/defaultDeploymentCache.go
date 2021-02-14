package deploy

import "sync"

type DefaultDeploymentCache struct {
}

func init() {
	deploymentCache = new(sync.Map)
}
func (defaultDeploymentCache DefaultDeploymentCache) Get(id string) ProcessDefinitionCacheEntry {
	load, ok := deploymentCache.Load(id)
	if ok {
		processDefinitionCacheEntry := load.(ProcessDefinitionCacheEntry)
		return processDefinitionCacheEntry
	}
	return ProcessDefinitionCacheEntry{}
}

func (defaultDeploymentCache DefaultDeploymentCache) Add(id string, object ProcessDefinitionCacheEntry) {
	deploymentCache.LoadOrStore(id, object)
}

func (defaultDeploymentCache DefaultDeploymentCache) Clear() {
	deploymentCache.Range(func(key, value interface{}) bool {
		deploymentCache.Delete(key)
		return true
	})
}
