package models

const DOMAIN = "ucn"

type User struct {
	UserId    int    `json:"user_id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	FullName  string `json:"full_name,omitempty"`
	CareerID  int    `json:"career_id,omitempty"`
	RUT       string `json:"RUT,omitempty"`
	Career    string `json:"career,omitempty"`
	Gender    string `json:"gender,omitempty"`
	GenderID  int    `json:"gender_id,omitempty"`
	IsAdmin   bool   `json:"is_admin"`
	Login
}

type UserResponse struct {
	Users []User `json:"users"`
}
