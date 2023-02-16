package models

import "time"

type Form struct {
	//user data
	StudentId   int    `json:"student_id,omitempty"`
	StudentName string `json:"student_full_name,omitempty"`
	CarrerName  string `json:"carrer_name,omitempty"`
	// Form table values
	FormId     int    `json:"form_id"`
	FormTitle  string `json:"form_title,omitempty"`
	FormDetail string `json:"form_detail,omitempty"`
	FormDate   string `json:"form_date,omitempty"`

	TotalSection   int               `json:"total_section,omitempty"`
	SectionContent []SectionContents `json:"section_content,omitempty"`
	Fields         []FieldData       `json:"fields,omitempty"`
	FieldsOrder    []FieldsOrder     `json:"fields_order,omitempty"`
}

type Forms struct {
	// Form table values
	FormId []Form `json:"forms,omitempty"`
}

type FormResponse struct {
	StudentId     int             `json:"student_id"`
	FormId        int             `json:"form_id"`
	FormTitle     int             `json:"form_title"`
	Date          time.Time       `json:"date"`
	SectionTime   []SectionTime   `json:"time_per_section"`
	FormResponses []FormResponses `json:"form_responses"`
}
type FormResponses struct {
	QuestionId             int    `json:"question_id,omitempty"`
	QuestionType           string `json:"question_type,omitempty"`
	AnswersItemId          int    `json:"answers_item_id,omitempty"`
	AnswersOptionId        int    `json:"answers_option_id,omitempty"`
	AnswersShortQuestionId int    `json:"answers_short_question_id,omitempty"`
	AnswersShortQuestion   string `json:"answers_short_question,omitempty"`
	AnswersSelectionId     int    `json:"answers_selection_id,omitempty"`
	AssigneScore           *int   `json:"assigne_score,omitempty"`
	Question               string `json:"question,omitempty"`
	IsOpenQuestion         bool   `json:"is_open_question"`
	Answer                 string `json:"answer,omitempty"`
	AnswerInt              int
	IsCorrect              *bool  `json:"is_correct,omitempty"`
	SectionId              int    `json:"section_id,omitempty"`
	SectionTitle           string `json:"section_title,omitempty"`
	ScoreForEachQuestion   int    `json:"score_for_each_question,omitempty"`
}

type SectionTime struct {
	SectionID   int `json:"section_id"`
	SectionTime int `json:"section_time"`
}

type AssigneScore struct {
	QuestionId   int `json:"question_id,omitempty"`
	StudentId    int `json:"student_id"`
	FormId       int `json:"form_id"`
	AssigneScore int `json:"assigne_score,omitempty"`
}

type AssigneScores struct {
	AssigneScore []AssigneScore `json:"assigne_scores,omitempty"`
}
