package product

import (
	"encoding/json"
	"fmt"
	"go_ecommerce/domain"
	"go_ecommerce/utils"
	"net/http"
	"strconv"
)

type ReqUpdateProducts struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	pId, err := strconv.Atoi(productId)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid product Id")
		return
	}

	var req ReqCreateProducts

	// NewDecoder will decode the information from the Request body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req) // this will convert that json into go struct

	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	_, err = h.svc.Update(domain.Product{
		ID: pId,
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImageUrl:    req.ImgUrl,
	})

	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	utils.SendData(w, "Successfully Update done", http.StatusOK)

}
