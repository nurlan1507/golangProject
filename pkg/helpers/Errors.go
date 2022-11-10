package helpers

import "errors"

var ErrNoRecord = errors.New("model: no matching record found")
var ConstraintError = errors.New("Constraint error")
var ErrDuplicate = errors.New("user with email already exists")
var PasswordIncorrect = errors.New("password incorrect")
var NoSuchUser = errors.New("no such user found")

var NotAuthorized = errors.New("you are not authorized")
var TokenError = errors.New("server Error")
var ExpiredToken = errors.New("token is expired")
var ExpiredRefreshToken = errors.New("RefreshToken is expired")

var EmailError = errors.New("smtp error")
var ValidationError = errors.New("validation error")
