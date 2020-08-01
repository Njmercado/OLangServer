package controller

import (
	"log"
	"regexp"
)

const specialCharacters = `!|,|;|.|:|-|_|\{|\}|\[|\]|\+|\*|\?|¡|¿|#|\$|\&|\(|\)`

func IsValidPassword(password string) bool {

	if len(password) < 8 {
		return false
	}

	hasUpperCase, _ := regexp.Match(`[A-Z]`, []byte(password))
	if !hasUpperCase {
		return false
	}

	hasNumbers, _ := regexp.Match(`\d`, []byte(password))
	if !hasNumbers {
		return false
	}

	hasSpecialCharacters, _ := regexp.Match(specialCharacters, []byte(password))
	if !hasSpecialCharacters {
		log.Print("no special characters")
		return false
	}

	return true
}

func IsValidUsername(username string) bool {

	if len(username) < 6 {
		return false
	}

	hasUpperCase, _ := regexp.Match(`[A-Z]`, []byte(username))
	if !hasUpperCase {
		return false
	}

	hasNumbers, _ := regexp.Match(`\d`, []byte(username))
	if !hasNumbers {
		return false
	}

	return true
}
