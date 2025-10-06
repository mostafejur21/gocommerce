package product

import "go_ecommerce/domain"

func (svc *service) Create(prdct domain.Product) (*domain.Product, error) {
	return svc.productRepo.Create(prdct)
}
