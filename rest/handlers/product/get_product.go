package product

import (
	"go_ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id") // Extract the Product Id from the request path

	pId, err := strconv.Atoi(productId) // convert string id into int
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.svc.Get(pId)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server error")
		return
	}

	if product == nil {
		utils.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	utils.SendData(w, product, http.StatusOK)
}
