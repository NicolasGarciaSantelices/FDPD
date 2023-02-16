package form

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	common_function "FDPD-BACKEND/src/common/function"
	common_models "FDPD-BACKEND/src/common/models"
	"FDPD-BACKEND/src/controllers/form/models"
	form_sql "FDPD-BACKEND/src/controllers/form/sql"
	"FDPD-BACKEND/src/controllers/user/constant"
)

func GetForms(c *gin.Context, db *sql.DB) {
	var (
		response *models.Forms
		err      error
	)
	response = &models.Forms{}

	*response, _ = form_sql.GetForms(db)

	common_function.SendResponse(c, common_function.CreateResponse(response, len(response.FormId), err))
}

func GetSectionByFormId(c *gin.Context, db *sql.DB) {
	var (
		response *models.Sections
		err      error
	)
	response = &models.Sections{}

	formIDstr := c.Param("formID")
	formID, _ := strconv.Atoi(formIDstr)

	*response, _ = form_sql.GetSection(formID, db)

	common_function.SendResponse(c, common_function.CreateResponse(response, len(response.SectionsInForm), err))
}

func GetQuestionBySectionId(c *gin.Context, db *sql.DB) {
	var (
		response *models.FieldsData
		err      error
	)
	response = &models.FieldsData{}

	//section id
	sectionIDstr := c.Param("sectionID")
	sectionID, _ := strconv.Atoi(sectionIDstr)

	*response, _ = form_sql.GetQuestion(sectionID, db)

	common_function.SendResponse(c, common_function.CreateResponse(response, len(response.Fields), err))
}

func GetFormById(c *gin.Context, db *sql.DB) {
	var (
		form              *models.Form
		sections          *models.Sections
		questionBySection *models.FieldsData
		_                 error
	)
	form = &models.Form{}
	//sections
	sections = &models.Sections{}

	formIDstr := c.Param("formID")
	formID, _ := strconv.Atoi(formIDstr)

	*sections, _ = form_sql.GetSection(formID, db)

	questionBySection = &models.FieldsData{}
	for _, section := range sections.SectionsInForm {
		questions, _ := form_sql.GetQuestion(section.Id, db)
		questionBySection.Fields = append(questionBySection.Fields, questions.Fields...)
	}

	form.FormId = formID
	form.TotalSection = len(sections.SectionsInForm)
	form.SectionContent = sections.SectionsInForm
	form.Fields = questionBySection.Fields
	form.FieldsOrder, _ = form_sql.GetQuestionOrder(formID, db)
	common_function.SendResponse(c, common_function.CreateResponse(form, len(form.Fields), nil))
}

func InputAnswer(c *gin.Context, db *sql.DB) {
	var (
		answers  *models.FormResponse
		response *common_models.Response
		err      error
	)
	answers = &models.FormResponse{}
	if err = c.ShouldBindJSON(&answers); err == nil {
		if err = form_sql.InsertAnswers(*answers, db); err == nil {
			response = &common_models.Response{
				Status:   constant.InsertAndsSuccesStatus,
				Code:     http.StatusCreated,
				Messages: constant.InsertAns,
			}
		} else {
			response = &common_models.Response{
				Status:   constant.InsertAnsErrorStatus,
				Code:     http.StatusBadRequest,
				Messages: err.Error(),
			}
		}

	}
	common_function.SendResponse(c, *response)
}

func GetAnswersByUserID(c *gin.Context, db *sql.DB) {
	var (
		answers  *models.FormResponse
		response *common_models.Response
	)
	answers = &models.FormResponse{}

	userIDstr := c.Param("userID")
	userID, _ := strconv.Atoi(userIDstr)

	formIDstr := c.Param("formID")
	formID, _ := strconv.Atoi(formIDstr)

	if ans, err := form_sql.GetAnswers(*answers, db, userID, formID); err == nil {
		response = &common_models.Response{
			Status:   constant.SuccesStatus,
			Code:     http.StatusOK,
			Messages: constant.GetAnsSucces,
			Data:     ans,
		}
	} else {
		response = &common_models.Response{
			Status:   constant.InsertAnsErrorStatus,
			Code:     http.StatusBadRequest,
			Messages: err.Error(),
		}
	}

	common_function.SendResponse(c, *response)
}

func GetFormsAnswersByUserID(c *gin.Context, db *sql.DB) {
	var (
		answers  *models.FormResponse
		response *common_models.Response
	)
	answers = &models.FormResponse{}

	userIDstr := c.Param("userID")
	userID, _ := strconv.Atoi(userIDstr)

	if ans, err := form_sql.GetFormAns(*answers, db, userID); err == nil {
		response = &common_models.Response{
			Status:   constant.SuccesStatus,
			Code:     http.StatusOK,
			Messages: constant.GetAnsFormSucces,
			Data:     ans,
		}
	} else {
		response = &common_models.Response{
			Status:   constant.InsertAnsErrorStatus,
			Code:     http.StatusBadRequest,
			Messages: err.Error(),
		}
	}

	common_function.SendResponse(c, *response)
}

func GetFormsAnswersForms(c *gin.Context, db *sql.DB) {
	var (
		answers  *models.FormResponse
		response *common_models.Response
	)
	answers = &models.FormResponse{}

	if ans, err := form_sql.GetAllFormAns(*answers, db); err == nil {
		response = &common_models.Response{
			Status:   constant.SuccesStatus,
			Code:     http.StatusOK,
			Messages: constant.GetAnsFormSucces,
			Data:     ans,
		}
	} else {
		response = &common_models.Response{
			Status:   constant.InsertAnsErrorStatus,
			Code:     http.StatusBadRequest,
			Messages: err.Error(),
		}
	}

	common_function.SendResponse(c, *response)
}

func GetFormsAnswersByForms(c *gin.Context, db *sql.DB) {
	var (
		answers  *models.FormResponse
		response *common_models.Response
	)
	answers = &models.FormResponse{}
	formIDstr := c.Param("formID")
	formID, _ := strconv.Atoi(formIDstr)

	if ans, err := form_sql.GetAllFormAnsByFormId(*answers, db, formID); err == nil {
		response = &common_models.Response{
			Status:   constant.SuccesStatus,
			Code:     http.StatusOK,
			Messages: constant.GetAnsFormSucces,
			Data:     ans,
		}
	} else {
		response = &common_models.Response{
			Status:   constant.InsertAnsErrorStatus,
			Code:     http.StatusBadRequest,
			Messages: err.Error(),
		}
	}

	common_function.SendResponse(c, *response)
}

func AssigneScore(c *gin.Context, db *sql.DB) {
	var (
		answers  *models.AssigneScore
		response *common_models.Response
		err      error
	)

	if err = c.ShouldBindJSON(&answers); err == nil {
		if err = form_sql.InsertAssigneScore(*answers, db); err == nil {
			response = &common_models.Response{
				Status:   constant.InsertAndsSuccesStatus,
				Code:     http.StatusCreated,
				Messages: constant.InsertAns,
			}
		} else {
			response = &common_models.Response{
				Status:   constant.InsertAnsErrorStatus,
				Code:     http.StatusBadRequest,
				Messages: err.Error(),
			}
		}

	}
	common_function.SendResponse(c, *response)
}

func AssigneScores(c *gin.Context, db *sql.DB) {
	var (
		answers  *models.AssigneScores
		response *common_models.Response
		err      error
	)

	if err = c.ShouldBindJSON(&answers); err == nil {
		if len(answers.AssigneScore) > 0 {
			for _, assigneScore := range answers.AssigneScore {
				if err = form_sql.InsertAssigneScore(assigneScore, db); err == nil {
					continue
				} else {
					break
				}
			}
			if err == nil {
				response = &common_models.Response{
					Status:   constant.InsertAndsSuccesStatus,
					Code:     http.StatusCreated,
					Messages: constant.InsertAns,
				}
			}
		} else {
			response = &common_models.Response{
				Status:   constant.InsertAnsErrorStatus,
				Code:     http.StatusBadRequest,
				Messages: err.Error(),
			}
		}

	}
	common_function.SendResponse(c, *response)
}
