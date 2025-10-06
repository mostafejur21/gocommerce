package user

import (
	"encoding/json"
	"fmt"
	"go_ecommerce/domain"
	"go_ecommerce/utils"
	"net/http"
)

type ReqCreateUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateUser

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		utils.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	user, err := h.svc.Create(domain.User{FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Password: req.Password, IsShopOwner: req.IsShopOwner})

	if err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.SendData(w, user, http.StatusCreated)
}

