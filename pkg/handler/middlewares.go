package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"testApp/pkg/helpers"
)

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func (h *Handler) EnableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-type", "application/json")
		next(w, r)
	}
}

func (h *Handler) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := r.Cookie("AccessToken")
		if err != nil {
			w.WriteHeader(401)
			return
		}
		_, e := h.TokenService.VerifyToken(accessToken.Value)
		if e != nil {
			if errors.Is(e.Err, helpers.ExpiredToken) {
				userId, _ := strconv.Atoi(fmt.Sprint(e.Payload["Id"]))
				fmt.Println(userId)
				_, err := h.TokenService.GetRefreshToken(userId)
				if err != nil {
					if errors.Is(err, helpers.ExpiredRefreshToken) {
						w.WriteHeader(401)
						return
					}
				}
				//
				token, err := h.TokenService.RefreshAccessToken(e.Payload)
				if err != nil {
					w.WriteHeader(401)
					return
				}
				newCookie := &http.Cookie{
					Name:     "AccessToken",
					Value:    token,
					HttpOnly: true,
					MaxAge:   2592000,
				}
				http.SetCookie(w, newCookie)
			} else {
				w.WriteHeader(401)
				return
			}
		}
		next(w, r)
	}
}

func (h *Handler) IsAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := r.Cookie("AccessToken")
		if err != nil {
			http.Redirect(w, r, "/signUp", 303)
			return
		}
		claims, _ := h.TokenService.GetClaims(accessToken.Value)
		fmt.Println(claims)
		if claims["Role"] == "Admin" {
			fmt.Println("ETO ADMIN VALIM")
			next(w, r)
		} else {
			http.Redirect(w, r, "/signUp", http.StatusMethodNotAllowed)
		}
	}
}
