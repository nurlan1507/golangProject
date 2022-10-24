package handler

import (
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testApp/pkg/service"
)

func (h *Handler) AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		accessToken, err := r.Cookie("AccessToken")
		fmt.Println(r.Cookies())
		if err != nil {
			http.Redirect(w, r, "/signUp", 303)
			return
		}

		_, e := h.TokenService.VerifyToken(accessToken.Value)
		if e != nil {
			if errors.Is(e.Err, service.ExpiredToken) {
				userId, ok := e.Payload["Id"].(int)
				if !ok {
				}
				//checking if refreshToken not expired in Db if it is expired then a user should login again
				//else : everything is ok, we regenerate a	ccessToken and set it to cookies
				_, err := h.TokenService.GetRefreshToken(userId)
				if err != nil {
					if errors.Is(err, service.ExpiredRefreshToken) {
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
					MaxAge:   300,
				}
				http.SetCookie(w, newCookie)
			}
		}
		next(w, r, ps)
	}
}

//err := beeep.Alert("title", "message", "")
//if err != nil {
//	fmt.Println(err)
//}
