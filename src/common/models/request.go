package common_models

type Pagination struct {
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	Sort    string `json:"sort"`
	Include string `json:"include"`
}
