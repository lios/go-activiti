package data

type AbstractDataManager interface {
	GetTaskEntityManager() TaskDataManager

	GetDefineEntityManager() DefineDataManager

	GetVariableEntityManager() VariableDataManager

	GetIdentityLinkEntityManager() IdentityLinkDataManager

	GetHistoricActinstEntityManager() HistoricActinstDataManager

	GetHistoricTaskEntityManager() HistoricTaskDataManager

	GetHistoricProcessEntityManager() HistoricProcessDataManager
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
