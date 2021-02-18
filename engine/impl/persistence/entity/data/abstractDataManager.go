package data

import (
	"github.com/lios/go-activiti/db"
	"github.com/lios/go-activiti/model"
)

type AbstractDataManager struct {
	model.TableModel
}

func (dataManagers AbstractDataManager) Insert(data interface{}) error {
	err := db.DB().Create(data).Error
	if err == nil {
		return err
	}
	return err
}

func (dataManagers AbstractDataManager) FindById(id int64, data interface{}) error {
	err := db.DB().Where("id = ?", id).Find(data).Error
	if err == nil {
		return err
	}
	return err
}
func (dataManagers AbstractDataManager) Delete(id int64) error {
	tableName := dataManagers.GetTableName()
	err := db.DB().Where("id = ?", id).Table(tableName).Error
	if err == nil {
		return err
	}
	return err
}

//var dataManager DataManager

//func IninDataManager(manager DataManager)  {
//	dataManager = manager
//}
//func(abstractDataManager AbstractDataManager)Insert(data interface{}) error {
//	err := db.DB().Create(data).Error
//	if err == nil {
//		return err
//	}
//	return err
//}

//func(AbstractDataManager) GetTaskEntityManager() *TaskDataManager {
//	return dataManager.TaskDataManager
//}
//
//func (AbstractDataManager) GetDefineEntityManager() *DefineDataManager {
//	return dataManager.DefineDataManager
//}
//func(AbstractDataManager) GetVariableEntityManager() *VariableDataManager {
//	return dataManager.VariableDataManager
//}
//
//func(AbstractDataManager) GetIdentityLinkEntityManager() *IdentityLinkDataManager {
//	return dataManager.IdentityLinkDataManager
//}
//
//func(AbstractDataManager) GetHistoricActinstEntityManager() *HistoricActinstDataManager {
//	return dataManager.HistoricActinstDataManager
//}
//
//func(AbstractDataManager) GetHistoricTaskEntityManager() *HistoricTaskDataManager {
//	return dataManager.HistoricTaskDataManager
//}
//
//func(AbstractDataManager) GetHistoricProcessEntityManager() *HistoricProcessDataManager {
//	return dataManager.HistoricProcessDataManager
//}
