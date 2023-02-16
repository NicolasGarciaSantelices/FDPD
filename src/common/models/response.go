package common_models

type Response struct {
	Status   string      `json:"status,omitempty"`
	Code     int         `json:"code,omitempty"`
	Messages string      `json:"messages,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}
