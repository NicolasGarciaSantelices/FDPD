package models

type Career struct {
	Id              string `json:"career_id"`
	CareerName      string `json:"career_name"`
	CareerShortName string `json:"career_short_name"`
}

type CareerResponse struct {
	Careers []Career `json:"careers"`
}
