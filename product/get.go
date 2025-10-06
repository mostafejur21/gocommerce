package product

import "go_ecommerce/domain"

func (svc *service) Get(id int) (*domain.Product, error) {
	return svc.productRepo.Get(id)
}
