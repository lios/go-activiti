package engine

type RepositoryService interface {
	Deploy(name string, key string, bytes []byte)
}
