package data

import (
	"github.com/lios/go-activiti/db"
	. "github.com/lios/go-activiti/model"
)

var resourceDataManager ResourceDataManager

type ResourceDataManager struct {
	AbstractDataManager
}

func (ResourceDataManager) GetResourceDataManager() ResourceDataManager {
	return resourceDataManager
}

func (define ResourceDataManager) FindDeployedProcessDefinitionByDeploymentId(deployMentId int64) (Bytearry, error) {
	bytearries := Bytearry{}
	err := db.DB().Where("`deployment_id`=?", deployMentId).Order("version DESC", true).First(&bytearries).Error

	return bytearries, err
}
