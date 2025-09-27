package product

import (
	"go_ecommerce/repo"
	"go_ecommerce/rest/middlewares"
)

type Handler struct {
	middlewares *middlewares.Middleware
	productRepo repo.ProductRepo
}

func NewHandler(
	middlewares *middlewares.Middleware,
	productRepo repo.ProductRepo,
) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}
