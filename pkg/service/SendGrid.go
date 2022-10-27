package service

import (
	"errors"
	"net/smtp"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func SendMessage(message string, to []string) error {
	from := "211369@astanait.edu.kz"
	host := "smtp.office365.com"
	auth := LoginAuth(from, "Baitasnur1507")
	err := smtp.SendMail(host+":587", auth, from, to, []byte(message))
	if err != nil {
		return err
	}
	return nil
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