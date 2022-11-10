package models

import (
	"time"
)

type UserModel struct {
	Id             int    `json:"id"`
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	AccessToken    string
	RefreshToken   string
	Role           string `json:"role"`
	RepeatPassword string `json:"repeatPassword"`
}
type SignUpModel struct {
	Email          string `json:"email"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
}

type RefreshToken struct {
	Token   string
	Expires time.Time
}

type TeacherInvite struct {
	InvitationId int
	TeacherId    int
	Token        string
}
