package models

type Gender struct {
	Id         string `json:"gender_id"`
	GenderName string `json:"gender_name"`
}

type GenderResponse struct {
	Genders []Gender `json:"genders"`
}
