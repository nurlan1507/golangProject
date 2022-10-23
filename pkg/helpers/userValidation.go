package helpers

import (
	"net/mail"
	"unicode"
)

// validation
func IsValidPassword(password string) bool {
	/*
	 * Password rules:
	 * at least 7 letters
	 * at least 1 number
	 * at least 1 upper case
	 * at least 1 special character
	 */
	var numbers = 0
	var upperCase = 0
	var specChar = 0
	for _, v := range password {
		switch {
		case unicode.IsNumber(v):
			numbers++
		case unicode.IsUpper(v):
			upperCase++
		case unicode.IsPunct(v) || unicode.IsSymbol(v):
			specChar++
		}
	}
	if numbers == 0 || upperCase == 0 || specChar == 0 {
		return false
	}
	return true
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}

func IsValidUsername(username string) bool {
	for _, v := range username {
		switch {
		case unicode.IsPunct(v):
			return false
		}
	}
	if len(username) < 5 || len(username) > 21 {
		return false
	}
	return true
}

func ArePasswordsEqual(pass1 string, pass2 string) bool {
	if pass1 != pass2 {
		return false
	}
	return true
}
