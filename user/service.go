package user

import "go_ecommerce/domain"

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{
		userRepo: userRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	usr, err := svc.userRepo.Create(user)

	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, err
	}

	return usr, nil
}

func (svc *service) Find(email string, pass string) (*domain.User, error) {
	usr, err := svc.userRepo.Find(email, pass)

	if err != nil {
		return nil, err
	}

	if usr == nil {
		return nil, err
	}

	return usr, nil
}
