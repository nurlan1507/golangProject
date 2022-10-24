package middlewares

import (
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"testApp/pkg/service"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		accessToken, err := r.Cookie("AccessToken")
		fmt.Println(r.Cookies())
		if err != nil {
			http.Redirect(w, r, "/signUp", 303)
			return
		}
		jwt := service.NewJWTManager()
		data, e := jwt.VerifyToken(accessToken.Value)
		if e != nil {
			if errors.Is(e.Err, service.ExpiredToken) {
				token, err := jwt.RefreshAccessToken(e.Payload)
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
			return
		}
		fmt.Printf("%+v", data)
		next(w, r, ps)
	}
}

//err := beeep.Alert("title", "message", "")
//if err != nil {
//	fmt.Println(err)
//}
