package models

type Indicators struct {
	FormID   int             `json:"form_id"`
	FormName string          `json:"form_title"`
	Gender   GenderIndicator `json:"gender"`
	Carrer   CarrerIndicator `json:"carrer"`
}

type GenderIndicator struct {
	Total int `json:"total"`
	Men   int `json:"male"`
	Women int `json:"female"`
}

type CarrerIndicator struct {
	Total   int       `json:"total"`
	Carrers []Carrers `json:"carrer"`
}

type Carrers struct {
	CareerID int    `json:"career_id,omitempty"`
	Career   string `json:"career,omitempty"`
}
