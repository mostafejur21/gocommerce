package user

import (
	"go_ecommerce/domain"
	userHandler "go_ecommerce/rest/handlers/user"
)

// domain can communicate with DB, redis or other domain
// so, when any domain dependent into other domain, those
// domain's signature will be in this port file

/*
	DDD rules:

	1. parent will never access directly the child
	2. child can access parent or grand parent or great grand parent
*/

type Service interface {
	userHandler.Service
}
type UserRepo interface {
	Create(user domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}
