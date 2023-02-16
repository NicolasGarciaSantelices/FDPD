package user

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	common_function "FDPD-BACKEND/src/common/function"
	common_models "FDPD-BACKEND/src/common/models"
	"FDPD-BACKEND/src/controllers/user/constant"
	"FDPD-BACKEND/src/controllers/user/models"
	user_sql "FDPD-BACKEND/src/controllers/user/sql"
)

func ValidateUser(c *gin.Context, db *sql.DB) {
	var (
		loginRequest *models.Login
		user         models.User
		response     *common_models.Response
		err          error
	)
	if err = c.ShouldBindJSON(&loginRequest); err == nil {
		if loginRequest.ValidateDomain() {
			if user = user_sql.UserExist(*loginRequest, db); user.Password == loginRequest.Password {
				user.Login.Password = ""
				response = &common_models.Response{
					Status:   constant.AuthStatus,
					Code:     http.StatusAccepted,
					Messages: constant.AuthSuccesfuly,
					Data:     user,
				}
			} else {
				response = &common_models.Response{
					Status:   constant.AuthErrorStatus,
					Code:     http.StatusNotAcceptable,
					Messages: constant.AuthError,
				}
			}

		} else {
			response = &common_models.Response{
				Status:   constant.AuthErrorStatus,
				Code:     http.StatusNotAcceptable,
				Messages: constant.AuthError,
			}
		}
	}
	common_function.SendResponse(c, *response)
}

func CreateUser(c *gin.Context, db *sql.DB) {
	var (
		users    *models.UserResponse
		response *common_models.Response
		err      error
	)
	response = &common_models.Response{}
	if err = c.ShouldBindJSON(&users); err == nil {
		response.Data, err = CreateValidUsers(users, db)
	}
	common_function.SendResponse(c, common_function.CreateResponse(response, 1, err))

}

func GetUser(c *gin.Context, db *sql.DB) {
	var (
		response *models.UserResponse
		err      error
		id       int
	)
	response = &models.UserResponse{}

	*response, _ = user_sql.GetUsers(id, db)

	common_function.SendResponse(c, common_function.CreateResponse(response, len(response.Users), err))
}

func UpdateUser(c *gin.Context, db *sql.DB) {
	var (
		updateRequest *models.User
		response      *common_models.Response
		err           error
	)

	if err = c.ShouldBindJSON(&updateRequest); err == nil {
		if err = user_sql.UpdateUserInfo(*updateRequest, db); err == nil {
			response = &common_models.Response{
				Messages: constant.UpdateSuccesfuly,
			}
		} else {
			response = &common_models.Response{
				Messages: err.Error(),
			}
		}

	}
	common_function.SendResponse(c, common_function.CreateResponse(response, 1, err))
}

func UpdateUserPassword(c *gin.Context, db *sql.DB) {
	var (
		updateRequest *models.User
		response      *common_models.Response
		err           error
	)

	if err = c.ShouldBindJSON(&updateRequest); err == nil {
		if err = user_sql.UpdateUserPassword(*updateRequest, db); err == nil {
			response = &common_models.Response{
				Messages: constant.UpdateSuccesfuly,
			}
		} else {
			response = &common_models.Response{
				Messages: err.Error(),
			}
		}

	}
	common_function.SendResponse(c, common_function.CreateResponse(response, 1, err))
}
