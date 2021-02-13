package data

import (
	"github.com/jinzhu/gorm"
	"github.com/lios/go-activiti/db"
	"github.com/lios/go-activiti/logger"
	. "github.com/lios/go-activiti/model"
	"time"
)

var deploymentDataManager DeploymentDataManager

type DeploymentDataManager struct {
	AbstractDataManager
	Deployment
}

func (DeploymentDataManager) GetDeploymentDataManager() DeploymentDataManager {
	return deploymentDataManager
}

//func (define DeploymentDataManager) FindById(id int64) (Deployment, error)  {
//	deployment := Deployment{}
//	if err :=db.DB().Where("id = ?",id).First(deployment).Error;err != nil {
//		logger.Error("find deploy by id err:",err)
//		return deployment, err
//	}
//	return deployment,nil
//}

func (define DeploymentDataManager) Deployments(name string, key string, bytes []byte) (err error) {
	deployment := Deployment{Name: name, Key: key, DeployTime: time.Now()}
	err = db.DB().Create(&deployment).Error
	if err != nil {
		logger.Error("Create deployment err", err)
		return err
	}
	defineManager := DefineDataManager{}
	bytearry, err := defineManager.FindDeployedProcessDefinitionByKey(key)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	var verion = 0
	verion = bytearry.Version
	verion++
	byte := Bytearry{Name: name, Key: key, Bytes: string(bytes), DeploymentId: deployment.Id, Version: verion}
	err = db.DB().Create(&byte).Error
	if err != nil {
		logger.Error("Create bytearry err", err)
		return err
	}
	return nil
}
