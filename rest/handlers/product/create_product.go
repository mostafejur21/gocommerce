package product

import (
	"encoding/json"
	"fmt"
	"go_ecommerce/domain"
	"go_ecommerce/utils"
	"net/http"
)

type ReqCreateProducts struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateProducts

	// NewDecoder will decode the information from the Request body
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req) // this will convert that json into go struct

	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	createProduct, err := h.svc.Create(domain.Product{
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

	utils.SendData(w, createProduct, http.StatusCreated)

}
