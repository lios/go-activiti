package engine

type RepositoryService interface {

	//流程部署
	Deploy(key string, name string, bytes string)
}
