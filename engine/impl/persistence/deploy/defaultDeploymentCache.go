package deploy

type DefaultDeploymentCache struct {
}

func (defaultDeploymentCache DefaultDeploymentCache) Get(id string) interface{} {
	return deploymentCache[id]
}

func (defaultDeploymentCache DefaultDeploymentCache) Add(id string, object interface{}) {
	deploymentCache[id] = object
}

func (defaultDeploymentCache DefaultDeploymentCache) Clear() {
	deploymentCache = make(map[string]interface{}, 0)
}
