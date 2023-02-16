package models

type OptionsFields struct {
	Id        string `json:"id"`
	Label     string `json:"label"`
	IsCorrect bool   `json:"is_correct,omitempty"`
	Custom    bool   `json:"custom"`
}

type Options struct {
	Option []OptionsFields `json:"options"`
}
