package models

import (
	"time"
)

type UserModel struct {
	Id          int
	Email       string
	Username    string
	Password    string
	AccessToken string
	Role        string
}
type RefreshToken struct {
	Token   string
	Expires time.Time
}
