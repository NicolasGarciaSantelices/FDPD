package sql

import (
	"database/sql"

	"FDPD-BACKEND/src/controllers/info/models"
)

func GetCareer(db *sql.DB) (response models.CareerResponse, err error) {
	rows, err := db.Query(`SELECT ` +
		`s.id, s.name, s.short_name ` +
		`FROM public.career s `)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var career models.Career
		err = rows.Scan(
			&career.Id,
			&career.CareerName,
			&career.CareerShortName,
		)
		if err == nil {
			response.Careers = append(response.Careers, career)
		}
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}

func GetGender(db *sql.DB) (response models.GenderResponse, err error) {
	rows, err := db.Query(`SELECT ` +
		`s.id, s.gender_name ` +
		`FROM public.gender s `)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var career models.Gender
		err = rows.Scan(
			&career.Id,
			&career.GenderName,
		)
		if err == nil {
			response.Genders = append(response.Genders, career)
		}
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}
