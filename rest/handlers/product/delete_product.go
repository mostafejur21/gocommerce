package product

import (
	"fmt"
	"go_ecommerce/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid Product Id")
		return
	}

	err = h.svc.Delete(pId)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	utils.SendData(w,"Successfully deleted", http.StatusOK)
}
