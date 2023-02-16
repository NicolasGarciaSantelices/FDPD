package models

type FieldData struct {
	Label        string          `json:"label"`
	Type         string          `json:"type"`
	Id           int             `json:"id"`
	Section      int             `json:"section,omitempty"`
	SubSection   string          `json:"sub_section,omitempty"`
	SubSectionID int             `json:"sub_section_id,omitempty"`
	Legend       *Legend         `json:"legend,omitempty"`
	Options      []OptionsFields `json:"options,omitempty"`
	Required     bool            `json:"required"`
}

type FieldsOrder struct {
	Id       string `json:"id"`
	Position int    `json:"position"`
}

type FieldsData struct {
	Fields []FieldData `json:"question"`
}
