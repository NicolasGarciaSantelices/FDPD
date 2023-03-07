package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	form_controller "FDPD-BACKEND/src/controllers/form"
	info_controller "FDPD-BACKEND/src/controllers/info"
	user_controller "FDPD-BACKEND/src/controllers/user"
)

func EndpointGroup(Engine *gin.Engine, db *sql.DB) error {

	api := Engine.Group("/v1")
	{

		info := api.Group("/info")
		{
			info.GET("/career", func(c *gin.Context) {
				info_controller.GetCareer(c, db)
			})
			info.GET("/gender", func(c *gin.Context) {
				info_controller.GetGender(c, db)
			})
		}

		user := api.Group("/user")
		{
			user.GET("/get", func(c *gin.Context) {
				user_controller.GetUser(c, db)
			})
			user.POST("/create", func(c *gin.Context) {
				user_controller.CreateUser(c, db)
			})
			user.POST("/auth", func(c *gin.Context) {
				user_controller.ValidateUser(c, db)
			})
			user.PUT("/update", func(c *gin.Context) {
				user_controller.UpdateUser(c, db)
			})
			user.PUT("/update/password", func(c *gin.Context) {
				user_controller.UpdateUserPassword(c, db)
			})

		}

		form := api.Group("/form")
		{
			form.GET("/get", func(c *gin.Context) {
				form_controller.GetForms(c, db)
			})

			form.GET("/:formID", func(c *gin.Context) {
				form_controller.GetFormById(c, db)
			})

			form.GET("/section/:formID", func(c *gin.Context) {
				form_controller.GetSectionByFormId(c, db)
			})

			form.GET("/questions/:sectionID", func(c *gin.Context) {
				form_controller.GetQuestionBySectionId(c, db)
			})
		}
		ans := api.Group("/answers")
		{

			ans.GET("", func(c *gin.Context) {
				form_controller.GetFormsAnswersForms(c, db)
			})
			ans.GET("/form/:formID", func(c *gin.Context) {
				form_controller.GetFormsAnswersByForms(c, db)
			})
			ans.POST("/forms", func(c *gin.Context) {
				form_controller.GetAnswersByUsersID(c, db)
			})
			ans.POST("/input", func(c *gin.Context) {
				form_controller.InputAnswer(c, db)
			})

			ans.GET("/user/:userID", func(c *gin.Context) {
				form_controller.GetFormsAnswersByUserID(c, db)
			})

			ans.GET("/user/:userID/:formID", func(c *gin.Context) {
				form_controller.GetAnswersByUserID(c, db)
			})
			ans.POST("/assigne-score", func(c *gin.Context) {
				form_controller.AssigneScore(c, db)
			})
			ans.POST("/assigne-scores", func(c *gin.Context) {
				form_controller.AssigneScores(c, db)
			})

		}

		ind := api.Group("/indicators")
		{
			ind.GET("", func(c *gin.Context) {
				form_controller.GetIndicators(c, db)
			})
		}

	}

	return nil
}
