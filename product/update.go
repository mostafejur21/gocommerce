package product

import "go_ecommerce/domain"

func (svc *service) Update(prdct domain.Product) (*domain.Product, error) {
	return svc.productRepo.Update(prdct)
}
