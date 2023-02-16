package models

import "strings"

type Login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (user *Login) ValidateDomain() bool {
	return strings.Contains(user.Email, DOMAIN)
}
