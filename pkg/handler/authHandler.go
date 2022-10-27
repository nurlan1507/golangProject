package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
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
	http.Redirect(w, r, "/home", http.StatusSeeOther)
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

func (h *Handler) SendEmail(w http.ResponseWriter, r *http.Request) {

	sendMessage("LOLOOLOLOLO", []string{"baitassovnurlan1507@gmail.com"})
	fmt.Println("Email sent successfully")
}

func sendMessage(message string, to []string) error {
	// user we are authorizing as
	from := "211369@astanait.edu.kz"
	// server we are authorized to send email through
	host := "smtp.office365.com"
	auth := LoginAuth(from, "Baitasnur1507")
	err := smtp.SendMail(host+":587", auth, from, to, []byte(message))
	if err != nil {
		return err
	}
	return nil
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unkown fromServer")
		}
	}
	return nil, nil
}
