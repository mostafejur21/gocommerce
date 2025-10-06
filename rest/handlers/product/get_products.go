package product

import (
	"go_ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.svc.Get(pId)
	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	if product == nil {
		utils.SendError(w, http.StatusNotFound, "Product Not Found")
		return
	}

	utils.SendData(w, product, http.StatusOK)
}
