package product

import "go_ecommerce/domain"

func (svc *service) List(page, limit int64) ([]*domain.Product, error) {
	return svc.productRepo.List(page, limit)
}
