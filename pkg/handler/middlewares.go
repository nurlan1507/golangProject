package handler

import (
	"errors"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testApp/pkg/helpers"
)

func (h *Handler) AuthMiddleware(next http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		accessToken, err := r.Cookie("AccessToken")
		if err != nil {
			http.Redirect(w, r, "/signUp", 303)
			return
		}
		_, e := h.TokenService.VerifyToken(accessToken.Value)
		if e != nil {
			if errors.Is(e.Err, helpers.ExpiredToken) {
				userId, ok := e.Payload["Id"].(int)
				if !ok {
				}
				//checking if refreshToken not expired in Db if it is expired then a user should login again
				//else : everything is ok, we regenerate a  accessToken and set it to cookies
				_, err := h.TokenService.GetRefreshToken(userId)
				if err != nil {
					if errors.Is(err, helpers.ExpiredRefreshToken) {
						http.Redirect(w, r, "/signUp", 400)
						return
					}
				}
				//
				token, err := h.TokenService.RefreshAccessToken(e.Payload)
				if err != nil {
					http.Error(w, "internal server error", 500)
				}
				newCookie := &http.Cookie{
					Name:     "AccessToken",
					Value:    token,
					HttpOnly: true,
					MaxAge:   2592000,
				}
				http.SetCookie(w, newCookie)
			} else {
				http.Redirect(w, r, "/signUp", 303)
				return
			}
		}
		next(w, r)
	}
}

func (h *Handler) IsAdmin(next http.HandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		accessToken, err := r.Cookie("AccessToken")
		if err != nil {
			http.Redirect(w, r, "/signUp", 303)
			return
		}
		claims, _ := h.TokenService.GetClaims(accessToken.Value)
		if claims["Role"] == "Admin" {
			next(w, r)
		} else {
			http.Redirect(w, r, "/signUp", http.StatusMethodNotAllowed)
		}
	}
}
