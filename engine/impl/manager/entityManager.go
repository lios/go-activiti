package manager

import (
	"github.com/lios/go-activiti/engine/impl/persistence/entity"
)

var entityManager EntityManager

type EntityManager struct {
	//DeploymentManager     *deploy.DeploymentManager
	ExecutionEntityManager *entity.ExecutionEntityManagerImpl
}

func init() {
	entityManager = EntityManager{}
	initEntityManagers()
}

func initEntityManagers() {
	//if entityManager.DeploymentManager == nil {
	//	entityManager.DeploymentManager = &deploy.DeploymentManager{}
	//}
	if entityManager.ExecutionEntityManager == nil {
		entityManager.ExecutionEntityManager = &entity.ExecutionEntityManagerImpl{}
	}

}

func GetEntityManager() EntityManager {
	return entityManager
}
