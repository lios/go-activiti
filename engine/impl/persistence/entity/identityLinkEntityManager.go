package entity

type IdentityLinkEntityManager interface {
	DeleteIdentityLinksByTaskId(taskId int64)
}
