package entity

import (
	"github.com/jinzhu/gorm"
	"github.com/lios/go-activiti/engine/impl/persistence/entity/data"
	"github.com/lios/go-activiti/logger"
	"github.com/lios/go-activiti/model"
)

var processInstanceDataManager data.DefineDataManager

type ProcessDefinitionEntityManagerImpl struct {
	AbstractEntityManager
}

func (processDefinitionEntityManager *ProcessDefinitionEntityManagerImpl) GetDataManager() data.DataManagers {
	return resourceDataManager
}

func (processDefinitionEntityManager *ProcessDefinitionEntityManagerImpl) FindProcessDefinitionById(processDefinitionId int64) ProcessDefinitionEntity {
	bytearry := model.Bytearry{}
	defineEntityManager := processInstanceDataManager
	defineEntityManager.FindById(processDefinitionId, &bytearry)
	definitionEntityImpl := ProcessDefinitionEntityImpl{}
	definitionEntityImpl.SetName(bytearry.Name)
	definitionEntityImpl.SetId(bytearry.Id)
	definitionEntityImpl.SetDeploymentId(bytearry.DeploymentId)
	return &definitionEntityImpl
}
func (processDefinitionEntityManager *ProcessDefinitionEntityManagerImpl) FindLatestProcessDefinitionByKey(processDefinitionKey string) (ProcessDefinitionEntity, error) {
	defineEntityManager := processInstanceDataManager
	bytearry, err := defineEntityManager.FindDeployedProcessDefinitionByKey(processDefinitionKey)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("FindDeployedProcessDefinitionByKey err :", err)
		return nil, err
	}
	definitionEntityImpl := ProcessDefinitionEntityImpl{}
	definitionEntityImpl.SetId(bytearry.Id)
	definitionEntityImpl.SetName(bytearry.Name)
	return &definitionEntityImpl, nil
}
