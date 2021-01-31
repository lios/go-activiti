package deploy

var deploymentCache map[string]interface{}

type DeploymentCache interface {
	Get(id string) interface{}
	Add(id string, object interface{})
	Clear()
}
