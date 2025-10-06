package user

import (
	"encoding/json"
	"go_ecommerce/utils"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	user, err := h.svc.Find(req.Email, req.Password)
	if err != nil {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if user == nil {
		utils.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	acceessToken, err := utils.CreateJWT(h.cnf.JwtSecretKey, utils.PayLoad{
		Sub:       user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utils.SendData(w, acceessToken, http.StatusOK)

}
