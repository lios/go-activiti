package data

import (
	"github.com/lios/go-activiti/db"
	. "github.com/lios/go-activiti/model"
	"github.com/prometheus/common/log"
	"time"
)

type DeploymentDataManager struct {
	AbstractDataManager
}

func (define DeploymentDataManager) Deployment(name string, key string, bytes []byte) (err error) {
	deployment := Deployment{Name: name, Key: key, DeployTime: time.Now()}
	err = db.DB().Create(&deployment).Error
	if err != nil {
		log.Infoln("Create deployment err", err)
		return err
	}
	defineManager := DefineDataManager{}
	bytearries, err := defineManager.FindDeployedProcessDefinitionByKey(key)
	if err != nil {
		return err
	}
	var verion = 0
	if bytearries != nil && len(bytearries) > 0 {
		verion = bytearries[0].Version
		verion++
	}
	byte := Bytearry{Name: name, Key: key, Bytes: string(bytes), DeploymentId: deployment.Id, Version: verion}
	err = db.DB().Create(&byte).Error
	if err != nil {
		log.Infoln("Create bytearry err", err)
		return err
	}
	return nil
}
