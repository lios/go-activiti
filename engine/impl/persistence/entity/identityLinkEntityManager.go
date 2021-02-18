package entity

import . "github.com/lios/go-activiti/model"

type IdentityLinkEntityManager interface {
	DeleteIdentityLinksByTaskId(taskId int64)

	CreateIdentityLink(identityLink IdentityLink) (err error)
}
