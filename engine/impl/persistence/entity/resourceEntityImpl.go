package entity

type ResourceEntityImpl struct {
	Name         string
	Bytes        []byte
	DeploymentId int64
}

func (resourceEntity *ResourceEntityImpl) GetName() string {
	return resourceEntity.Name
}

func (resourceEntity *ResourceEntityImpl) SetName(name string) {
	resourceEntity.Name = name
}

func (resourceEntity *ResourceEntityImpl) GetBytes() []byte {
	return resourceEntity.Bytes
}

func (resourceEntity *ResourceEntityImpl) SetBytes(bytes []byte) {

}

func (resourceEntity *ResourceEntityImpl) GetDeploymentId() int64 {
	return resourceEntity.DeploymentId
}

func (resourceEntity *ResourceEntityImpl) SetDeploymentId(deploymentId int64) {

}
