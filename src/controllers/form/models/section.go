package models

type SectionContents struct {
	Id    int    `json:"id,omitempty"`
	Title string `json:"title"`
	Score int    `json:"score_for_each_question,omitempty"`
	Order string `json:"order"`
}

type Sections struct {
	SectionsInForm []SectionContents `json:"sections_in_form"`
}
