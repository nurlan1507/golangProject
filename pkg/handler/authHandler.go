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
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	GroupId   int    `json:"groupId"`
	Validator *helpers.Validation
}

func (h *Handler) SignUpPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := r.ParseForm()

	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
		log.Fatal(err)
	}
	var newUser = &models.SignUpModel{}
	err = json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
		return
	}

	AuthForm := &AuthForm{
		Email:     newUser.Email,
		Username:  newUser.Username,
		Password:  newUser.Password,
		GroupId:   newUser.GroupId,
		Validator: helpers.NewValidation(),
	}
	AuthForm.Validator.Check(helpers.IsValidEmail(AuthForm.Email), "email", "email is not valid")
	AuthForm.Validator.Check(helpers.IsValidUsername(AuthForm.Username), "username", "username should not contain [.!?\\-] and be less than 5 symbols")
	AuthForm.Validator.Check(helpers.IsValidPassword(AuthForm.Password), "password", "Password rules: at least 7 letters \n at least 1 number \n at least 1 upper case \n at least 1 special character")
	AuthForm.Validator.Check(helpers.ArePasswordsEqual(AuthForm.Password, newUser.RepeatPassword), "repeatPassword", "passwords do not match")
	if AuthForm.Validator.Valid() == false {
		res, err := json.Marshal(AuthForm.Validator.Errors)
		h.Loggers.ErrorLogger.Println(err)
		w.WriteHeader(400)
		w.Write(res)
		return
	}
	user, err := h.UserService.SignUp(AuthForm.Email, AuthForm.Username, AuthForm.Password, AuthForm.GroupId)
	if err != nil {
		res, err := json.Marshal(AuthForm.Validator.Errors)
		if errors.Is(err, helpers.ErrDuplicate) {
			AuthForm.Validator.Errors["duplicate"] = err.Error()
			w.WriteHeader(400)
			if err != nil {
				return
			}
			w.Write(res)
			return
		}
		h.Loggers.ErrorLogger.Println(err)
		w.WriteHeader(500)
		//helpers.BadRequest(w, r, err)
		w.Write(res)
		return
	}
	AccessTokenCookie := &http.Cookie{
		Name:     "AccessToken",
		Value:    user.AccessToken,
		MaxAge:   300,
		HttpOnly: true,
		Secure:   true,
	}
	RefreshTokenCookie := &http.Cookie{
		Name:     "RefreshToken",
		Value:    user.RefreshToken,
		MaxAge:   2592000,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, AccessTokenCookie)
	http.SetCookie(w, RefreshTokenCookie)
	marshal, err := json.Marshal(&user)
	if err != nil {
		return
	}
	w.Write(marshal)
	return
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
	if err != nil {
		return
	}
	authForm := &AuthForm{}
	json.NewDecoder(r.Body).Decode(&authForm)
	authForm.Validator = helpers.NewValidation()
	//validation
	authForm.Validator.Check(helpers.IsFulfilled(authForm.Email), "email", "email is empty")
	authForm.Validator.Check(helpers.IsFulfilled(authForm.Password), "password", "field password is empty")
	if authForm.Validator.Valid() == false {
		res, _ := json.Marshal(authForm.Validator.Errors)
		w.WriteHeader(400)
		w.Write(res)
		return
	}
	result, err := h.UserService.SignIn(authForm.Email, authForm.Password)
	if err != nil {
		w.WriteHeader(400)
		if errors.Is(err, helpers.NoSuchUser) {
			authForm.Validator.Errors["email"] = err.Error()
			res, _ := json.Marshal(authForm.Validator.Errors)
			w.Write(res)
			return
		}
		if errors.Is(err, helpers.PasswordIncorrect) {
			authForm.Validator.Errors["password"] = err.Error()
			res, _ := json.Marshal(authForm.Validator.Errors)
			w.Write(res)
			return
		}
		w.WriteHeader(500)
		return
	}
	cookie := &http.Cookie{
		Name:     "AccessToken",
		Value:    result.AccessToken,
		MaxAge:   300,
		HttpOnly: true,
		Secure:   true,
	}
	RefreshTokenCookie := &http.Cookie{
		Name:     "RefreshToken",
		Value:    result.RefreshToken,
		MaxAge:   2592000,
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, RefreshTokenCookie)
	http.SetCookie(w, cookie)
	response, _ := json.Marshal(result)
	w.Header().Add("accessT", result.AccessToken)
	w.WriteHeader(200)
	w.Write(response)
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

func (h *Handler) GetNewAccessToken(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := r.Cookie("RefreshToken")

	if err != nil {
		w.WriteHeader(401)
		h.Loggers.ErrorLogger.Println(err)
		return
	}
	claims, err := h.TokenService.GetClaims(refreshToken.Value)
	fmt.Println(claims)
	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
		w.WriteHeader(401)
		return
	}
	token, err := h.TokenService.RefreshAccessToken(claims)
	if err != nil {
		h.Loggers.ErrorLogger.Println(err)
		w.WriteHeader(500)
		return
	}
	newCookie := &http.Cookie{
		Name:   "AccessToken",
		Value:  token,
		MaxAge: 300,
	}
	http.SetCookie(w, newCookie)
	res, _ := json.Marshal(&newCookie)
	w.Write(res)
}
