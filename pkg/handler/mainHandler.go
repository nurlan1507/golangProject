package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"testApp/pkg/models"
	"testApp/pkg/templateData"
)

func (h *Handler) HOME(AuthDada models.UserModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.render(w, "home.tmpl", nil, 200)
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("AccessToken")
	if err != nil {
		return
	}
	claims, err := h.TokenService.GetClaims(cookie.Value)
	userID, _ := strconv.Atoi(fmt.Sprint(claims["Id"]))
	AuthData := &models.UserModel{
		Id:       userID,
		Email:    fmt.Sprint(claims["Email"]),
		Username: fmt.Sprint(claims["Username"]),
		Role:     fmt.Sprint(claims["Role"]),
	}
	fmt.Printf("%+v", AuthData)
	h.render(w, "home.tmpl", templateData.NewTemplateData(AuthData, nil), 200)
}
