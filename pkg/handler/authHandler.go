package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"testApp/pkg/helpers"
	"testApp/pkg/models"
	"testApp/pkg/templateData"
)

type AuthForm struct {
	Email     string
	Username  string
	Password  string
	Validator *helpers.Validation
}

func (h *Handler) SignUpPost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
		log.Fatal(err)
	}
	AuthForm := &AuthForm{
		Email:     r.PostForm.Get("email"),
		Username:  r.PostForm.Get("username"),
		Password:  r.PostForm.Get("password"),
		Validator: helpers.NewValidation(),
	}

	AuthForm.Validator.Check(helpers.IsValidEmail(AuthForm.Email), "email", "email is not valid")
	AuthForm.Validator.Check(helpers.IsValidUsername(AuthForm.Username), "username", "username should not contain [.!?\\-] and be less than 5 symbols")
	AuthForm.Validator.Check(helpers.IsValidPassword(AuthForm.Password), "password", "Password rules: at least 7 letters \n at least 1 number \n at least 1 upper case \n at least 1 special character")
	AuthForm.Validator.Check(helpers.ArePasswordsEqual(AuthForm.Password, r.PostForm.Get("repeatPassword")), "repeatPassword", "passwords do not match")
	if AuthForm.Validator.Valid() == false {
		h.render(w, "signUp.tmpl", templateData.NewTemplateData(&models.UserModel{}, AuthForm))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	user, err := h.UserService.SignUp(AuthForm.Email, AuthForm.Username, AuthForm.Password)
	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
		helpers.BadRequest(w, r, err)
		return
	}
	cookie := &http.Cookie{
		Name:     "AccessToken",
		Value:    user.AccessToken,
		MaxAge:   300,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/home", http.StatusSeeOther)
	return
	//json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	up := h.UserService.GetUsers()
	json.NewEncoder(w).Encode(up)
}
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	AuthForm := &AuthForm{
		Email:     r.PostForm.Get("email"),
		Username:  r.PostForm.Get("username"),
		Password:  r.PostForm.Get("password"),
		Validator: helpers.NewValidation(),
	}
	newUser := &models.UserModel{
		Username:    "n",
		Email:       "SD",
		Password:    "SDS",
		AccessToken: "SSDA",
	}

	h.render(w, "signUp.tmpl", templateData.NewTemplateData(newUser, AuthForm))
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	h.render(w, "signIn.tmpl", nil)
}
