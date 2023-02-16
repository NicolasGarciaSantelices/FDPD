package models

type Legend struct {
	Id         int               `json:"-"`
	LabelFirst string            `json:"labelFirst"`
	Columns    []SectionContents `json:"columns,omitempty"`
}
