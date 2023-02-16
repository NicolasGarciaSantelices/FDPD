package info

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	common_function "FDPD-BACKEND/src/common/function"
	"FDPD-BACKEND/src/controllers/info/models"
	info_sql "FDPD-BACKEND/src/controllers/info/sql"
)

func GetCareer(c *gin.Context, db *sql.DB) {
	var (
		response *models.CareerResponse
		err      error
	)
	response = &models.CareerResponse{}

	*response, _ = info_sql.GetCareer(db)

	common_function.SendResponse(c, common_function.CreateResponse(response, len(response.Careers), err))
}

func GetGender(c *gin.Context, db *sql.DB) {
	var (
		response *models.GenderResponse
		err      error
	)
	response = &models.GenderResponse{}

	*response, _ = info_sql.GetGender(db)

	common_function.SendResponse(c, common_function.CreateResponse(response, len(response.Genders), err))
}
