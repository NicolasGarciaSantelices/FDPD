package models

type ResponseCreateUser struct {
	UserCreated   int                 `json:"user_created"`
	UserError     int                 `json:"user_not_created"`
	UserErrorInfo []UserWithErrorInfo `json:"users_with_error"`
	Messages      string              `json:"messages,omitempty"`
	Data          interface{}         `json:"data,omitempty"`
}

type UserWithErrorInfo struct {
	User  User   `json:"user"`
	Error string `json:"error_info"`
}
