package info

import (
	"database/sql"

	"FDPD-BACKEND/src/controllers/user/constant"
	"FDPD-BACKEND/src/controllers/user/models"
	user_sql "FDPD-BACKEND/src/controllers/user/sql"
)

func CreateValidUsers(users *models.UserResponse, db *sql.DB) (response models.ResponseCreateUser, err error) {

	response = models.ResponseCreateUser{}
	for _, user := range users.Users {
		if user.ValidateDomain() {
			if err = user_sql.CreateStudent(user, db); err == nil {
				response.UserCreated++
			} else {
				response.UserError++

				userWithError := models.UserWithErrorInfo{
					User:  user,
					Error: constant.ErrorUserExist.Error(),
				}
				response.UserErrorInfo = append(response.UserErrorInfo, userWithError)
			}
		} else {
			response.UserError++

			userWithError := models.UserWithErrorInfo{
				User:  user,
				Error: constant.ErrorInAuth.Error(),
			}
			response.UserErrorInfo = append(response.UserErrorInfo, userWithError)
		}
	}
	return
}
