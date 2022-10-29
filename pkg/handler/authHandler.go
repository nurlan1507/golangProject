package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
		h.render(w, "signUp.tmpl", templateData.NewTemplateData(&models.UserModel{}, AuthForm), 400)
		return
	}
	user, err := h.UserService.SignUp(AuthForm.Email, AuthForm.Username, AuthForm.Password)
	if err != nil {
		if errors.Is(err, helpers.ErrDuplicate) {
			AuthForm.Validator.Errors["duplicate"] = err.Error()
			h.render(w, "signUp.tmpl", templateData.NewTemplateData(&models.UserModel{}, AuthForm), 400)
			return
		}
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
	http.Redirect(w, r, "/", http.StatusSeeOther)
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
	h.render(w, "signUp.tmpl", templateData.NewTemplateData(nil, AuthForm), 200)
}

func (h *Handler) SignInPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	AuthForm := &AuthForm{Email: r.PostForm.Get("email"), Password: r.PostForm.Get("password"), Validator: helpers.NewValidation()}
	data := templateData.NewTemplateData(nil, AuthForm)
	res, err := h.UserService.SignIn(r.PostForm.Get("email"), r.PostForm.Get("password"))
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, helpers.NoSuchUser) {
			AuthForm.Validator.Errors["NotFound"] = err.Error()
		}
		if errors.Is(err, helpers.PasswordIncorrect) {
			AuthForm.Validator.Errors["PasswordNotMatch"] = err.Error()
		}
		data.Form = AuthForm
		h.render(w, "signIn.tmpl", data, http.StatusBadRequest)
		return
	}
	data.AuthData = res
	cookie := &http.Cookie{
		Name:     "AccessToken",
		Value:    res.AccessToken,
		MaxAge:   300,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	AuthForm := &AuthForm{
		Email:     r.PostForm.Get("email"),
		Username:  r.PostForm.Get("username"),
		Password:  r.PostForm.Get("password"),
		Validator: helpers.NewValidation(),
	}
	h.render(w, "signIn.tmpl", templateData.NewTemplateData(nil, AuthForm), 200)

	return
}

func (h *Handler) SignUpTeacher(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	claims, err := h.TokenService.GetClaims(token)
	if err != nil {
		return
	}
	AuthForm := &AuthForm{
		Email:     fmt.Sprint(claims["Email"]),
		Username:  fmt.Sprint(claims["Username"]),
		Validator: helpers.NewValidation(),
	}
	userIdstr := fmt.Sprint(claims["Id"])
	Id, _ := strconv.Atoi(userIdstr)
	h.render(w, "signUpTeacher.tmpl", templateData.NewTemplateData(&models.UserModel{Id: Id}, AuthForm), 200)
}

func (h *Handler) SignUpTeacherPost(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("userId")
	fmt.Println("ASDASDASDAS")
	err := r.ParseForm()
	if err != nil {
		return
	}
	AuthForm := &AuthForm{
		Password:  r.PostForm.Get("password"),
		Validator: helpers.NewValidation(),
	}
	AuthForm.Validator.Check(helpers.IsValidPassword(AuthForm.Password), "password", "Password rules: at least 7 letters \n at least 1 number \n at least 1 upper case \n at least 1 special character")
	AuthForm.Validator.Check(helpers.ArePasswordsEqual(AuthForm.Password, r.PostForm.Get("repeatPassword")), "repeatPassword", "passwords do not match")
	if !AuthForm.Validator.Valid() {
		h.render(w, "signUpTeacher.tmpl", templateData.NewTemplateData(nil, AuthForm), 200)
		return
	}
	idInt, _ := strconv.Atoi(id)
	_, err = h.UserService.SignUpTeacher(idInt, AuthForm.Password)
	if err != nil {
		fmt.Println(err)
		return
	}
	http.Redirect(w, r, "/home", 303)
}

func (h *Handler) GetUserData(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("AccessToken")
	if err != nil {
		http.Redirect(w, r, "/signUp", 303)
		return
	}

	claims, err := h.TokenService.GetClaims(cookie.Value)
	if err != nil {
		http.Error(w, "Internal server Error", 500)
		return
	}
	id, _ := strconv.Atoi(fmt.Sprint(claims["id"]))
	userData := &models.UserModel{
		Id:       id,
		Email:    fmt.Sprint(claims["Email"]),
		Username: fmt.Sprint(claims["Username"]),
		Role:     fmt.Sprint(claims["Role"]),
	}
	json.Marshal(userData)

}
