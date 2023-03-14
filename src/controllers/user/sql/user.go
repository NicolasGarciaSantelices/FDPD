package sql

import (
	"database/sql"
	"strings"

	"FDPD-BACKEND/src/controllers/user/models"
	"FDPD-BACKEND/src/utils"
)

func CreateStudent(users models.User, db *sql.DB) (err error) {
	users.RUT = strings.Replace(users.RUT, ".", "", -3)
	users.RUT = strings.Replace(users.RUT, "-", "", -1)
	password := users.RUT[len(users.RUT)-4:]
	insertDynStmt :=
		`INSERT INTO public.student
	(first_name, last_name, full_name, career_id, rut, gender_id, email, "password")
	VALUES($1, $2, $3, $4, $5, $6, $7, $8)`

	_, e := db.Exec(insertDynStmt, users.FirstName, users.LastName, users.FullName, users.CareerID, users.RUT, users.GenderID, users.Email, password)
	if e != nil {
		utils.RecoverError()
		return e
	}

	return nil
}

func GetUsers(careerId int, db *sql.DB) (response models.UserResponse, err error) {
	rows, err := db.Query(`SELECT ` +
		`s.id, s.first_name, s.last_name, s.full_name, s.career_id, c.name, s.rut, s.gender_id,g.gender_name, s.email ` +
		`FROM public.student s ` +
		`INNER JOIN public.career c ` +
		`ON c.id = s.career_id ` +
		`INNER JOIN public.gender g ` +
		`ON g.id = s.gender_id `)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.UserId,
			&user.FirstName,
			&user.LastName,
			&user.FullName,
			&user.CareerID,
			&user.Career,
			&user.RUT,
			&user.GenderID,
			&user.Gender,
			&user.Email,
		)
		if err == nil {
			response.Users = append(response.Users, user)
		}
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}

func GetUsersByID(user models.User, db *sql.DB) (response models.UserResponse, err error) {
	rows, err := db.Query(`SELECT `+
		`s.id, s.first_name, s.last_name, s.full_name, s.career_id, c.name, s.rut, s.gender_id,g.gender_name, s.email, s.password `+
		`FROM public.student s `+
		`INNER JOIN public.career c `+
		`ON c.id = s.career_id `+
		`INNER JOIN public.gender g `+
		`ON g.id = s.gender_id `+
		`WHERE s.id =$1`, user.UserId)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var user models.User
		err = rows.Scan(
			&user.UserId,
			&user.FirstName,
			&user.LastName,
			&user.FullName,
			&user.CareerID,
			&user.Career,
			&user.RUT,
			&user.GenderID,
			&user.Gender,
			&user.Email,
			&user.Password,
		)
		if err == nil {
			response.Users = append(response.Users, user)
		}
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}

func UserExist(login models.Login, db *sql.DB) (user models.User) {

	adminList := make(map[string]bool)
	rows, err := db.Query(`SELECT ` +
		`s.email ` +
		`FROM public.admin a ` +
		`INNER JOIN public.student s ON s.id = a.user_id `)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var email string
		err = rows.Scan(
			&email,
		)
		if err == nil {
			adminList[email] = true
		}
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	rows, err = db.Query(`SELECT `+
		`s.id, s.first_name, s.last_name, s.full_name, s.career_id, c.name, s.rut, s.gender_id,g.gender_name, s.email, s.password `+
		`FROM public.student s `+
		`INNER JOIN public.career c `+
		`ON c.id = s.career_id `+
		`INNER JOIN public.gender g `+
		`ON g.id = s.gender_id `+
		`WHERE s.email =$1`, login.Email)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&user.UserId,
			&user.FirstName,
			&user.LastName,
			&user.FullName,
			&user.CareerID,
			&user.Career,
			&user.RUT,
			&user.GenderID,
			&user.Gender,
			&user.Email,
			&user.Password,
		)
		if err == nil {
			if adminList[user.Email] {
				user.IsAdmin = true
			}
			return
		}
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}
	return
}

func UpdateUserInfo(userUpdate models.User, db *sql.DB) (err error) {

	insertDynStmt :=
		`UPDATE public.student SET ` +
			`first_name = $1, ` +
			`last_name = $2, ` +
			`full_name = $3, ` +
			`career_id = $4, ` +
			`rut = $5, ` +
			`gender_id = $6, ` +
			`email = $7 ` +
			`WHERE id = $8 `

	_, e := db.Exec(
		insertDynStmt,
		userUpdate.FirstName,
		userUpdate.LastName,
		userUpdate.FullName,
		userUpdate.CareerID,
		userUpdate.RUT,
		userUpdate.GenderID,
		userUpdate.Email,
		userUpdate.UserId)
	if e != nil {
		utils.RecoverError()
		return e
	}
	return nil
}

func UpdateUserPassword(userUpdate models.User, db *sql.DB) (err error) {
	insertDynStmt :=
		`UPDATE public.student SET ` +
			`password = $1 ` +
			`WHERE email = $2 `

	_, e := db.Exec(
		insertDynStmt,
		userUpdate.Password,
		userUpdate.Email)
	if e != nil {
		utils.RecoverError()
		return e
	}
	return nil
}

// TODO:funcion para obtener todos usuarios en la bd por carrera , retornar error si existe
func GetUsersByCareerName(careerName string) (response []models.UserResponse, err error) {

	return nil, nil
}
