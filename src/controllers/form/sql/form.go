package sql

import (
	"database/sql"
	"strconv"
	"time"

	"FDPD-BACKEND/src/controllers/form/models"
	"FDPD-BACKEND/src/utils"
)

func GetForms(db *sql.DB) (response models.Forms, err error) {
	rows, err := db.Query(`SELECT ` +
		`id, title, description ` +
		`FROM public.form `)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var form models.Form
		err = rows.Scan(
			&form.FormId,
			&form.FormTitle,
			&form.FormDetail,
		)
		if err == nil {
			response.FormId = append(response.FormId, form)
		}
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}

func GetSection(formID int, db *sql.DB) (response models.Sections, err error) {
	rows, err := db.Query(`SELECT `+
		`s.id,s.title, s.score_for_each_question,fs.order `+
		`FROM public.form_sections fs `+
		`INNER JOIN public.form f ON f.id = fs.form_id `+
		`INNER JOIN public.section s ON s.id = fs.section_id `+
		`WHERE fs.form_id = $1 `, formID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var section models.SectionContents
		err = rows.Scan(
			&section.Id,
			&section.Title,
			&section.Score,
			&section.Order,
		)
		if err == nil {
			response.SectionsInForm = append(response.SectionsInForm, section)
		}
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}

func GetQuestion(sectionID int, db *sql.DB) (response models.FieldsData, err error) {
	rows, err := db.Query(`SELECT `+
		`COALESCE(q.description,''), qt.type, q.id,COALESCE(q.image_url,''),COALESCE(q.title,''),COALESCE(q.question_description,''), q.input_text,COALESCE(q.img_titles,'') `+
		`FROM public.section_questions sq `+
		`INNER JOIN public.section s ON s.id = sq.section_id `+
		`INNER JOIN public.question q ON q.id = sq.question_id `+
		`INNER JOIN public.question_type qt ON q.type_id = qt.id `+
		`WHERE s.id = $1`, sectionID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var question models.FieldData
		err = rows.Scan(
			&question.Label,
			&question.Type,
			&question.Id,
			&question.ImgURL,
			&question.Title,
			&question.QuestionDescription,
			&question.InputText,
			&question.ImgTitle,
		)
		if err == nil {
			if *question.ImgURL == "" {
				question.ImgURL = nil
			}
			if *question.Title == "" {
				question.Title = nil
			}
			if *question.QuestionDescription == "" {
				question.QuestionDescription = nil
			}
			question.Section = sectionID
			question.Required = true

			switch question.Type {
			case "RADIO":
				question.Options, _ = GetRadioOptions(question.Id, db)
			case "LINEAR":
				question.Options, _ = GetLinearOptions(question.Id, db)
				question.Legend, question.SubSection, question.SubSectionID, _ = GetLinearLegend(question.Id, db)
			case "SHORT_ANSWER":
				// no se hace nada ya que no se tienen respuestas predeterminadas de este tipo
				break
			}

			response.Fields = append(response.Fields, question)
		}
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}

func GetRadioOptions(questionID int, db *sql.DB) (response []models.OptionsFields, err error) {
	rows, err := db.Query(`SELECT `+
		`qso.id,qso.option, qso.is_correct `+
		`FROM public.question_selection_option qso `+
		`INNER JOIN public.question q ON q.id = qso.question_id `+
		`WHERE q.id = $1`, questionID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var options models.OptionsFields
		err = rows.Scan(
			&options.Id,
			&options.Label,
			&options.IsCorrect,
		)
		if err == nil {
			options.Custom = false
			response = append(response, options)
		}
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}

func GetLinearLegend(QuestionId int, db *sql.DB) (*models.Legend, string, int, error) {
	var (
		subSection   string
		SubSectionID int
	)
	response := &models.Legend{}

	rows, err := db.Query(`SELECT `+
		`ql.id,ql.description,ss.title,ql.sub_section_id  `+
		`FROM public.question_linear ql `+
		`INNER JOIN public.sub_section ss ON ss.id = ql.sub_section_id `+
		`WHERE ql.question_id = $1 ORDER BY ql.id`, QuestionId)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var (
			legend models.Legend
		)

		err = rows.Scan(
			&legend.Id,
			&legend.LabelFirst,
			&subSection,
			&SubSectionID,
		)
		if err == nil {
			//legends
			rowsColumns, err := db.Query(`SELECT `+
				`qli.id,qli.item `+
				`FROM public.question_linear_item qli `+
				`WHERE qli.question_id = $1`, legend.Id)
			if err != nil {
				panic(err)
			}
			defer rowsColumns.Close()
			for rowsColumns.Next() {
				var SectionContents models.SectionContents
				err = rowsColumns.Scan(
					&SectionContents.Id,
					&SectionContents.Title,
				)
				if err == nil {
					legend.Columns = append(legend.Columns, SectionContents)
				}
			}

			//options
		}
		response = &legend
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return response, subSection, SubSectionID, nil
}

func GetLinearOptions(questionID int, db *sql.DB) (response []models.OptionsFields, err error) {
	rows, err := db.Query(`SELECT `+
		`qso.id,qso.option,COALESCE(qso.image_url) `+
		`FROM public.question_linear_option qso `+
		`INNER JOIN public.question q ON q.id = qso.question_id `+
		`WHERE q.id = $1`, questionID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var options models.OptionsFields
		err = rows.Scan(
			&options.Id,
			&options.Label,
			&options.ImageURL,
		)
		if err == nil {
			options.Custom = false
			response = append(response, options)
		}
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}

func GetQuestionOrder(formID int, db *sql.DB) (response []models.FieldsOrder, err error) {
	rows, err := db.Query(`SELECT `+
		`qo.order, qo.question_id `+
		`FROM public.form_order qo `+
		`WHERE qo.form_id = $1`, formID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var fieldsOrder models.FieldsOrder
		err = rows.Scan(
			&fieldsOrder.Position,
			&fieldsOrder.Id,
		)
		if err == nil {
			response = append(response, fieldsOrder)
		}

	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}

func InsertAnswers(answers models.FormResponse, db *sql.DB) (err error) {

	//if exist prev answers for user and form, delete
	var prevID int
	//save answers by user
	selectDynStmt :=
		`SELECT id FROM public.form_answers_user ` +
			`WHERE student_id = $1 and  form_id = $2 `
	prevAnsRows, e := db.Query(
		selectDynStmt,
		answers.StudentId,
		answers.FormId,
	)

	if e != nil {
		utils.RecoverError()
		return e
	}
	defer prevAnsRows.Close()
	for prevAnsRows.Next() {
		err = prevAnsRows.Scan(
			&prevID,
		)
		if err != nil {
			return err
		}
	}

	if prevID != 0 {
		//delete times per section
		deleteDynStmt :=
			`DELETE FROM public.time_per_section ` +
				`WHERE "user" = $1 ` +
				``
		_, e := db.Query(
			deleteDynStmt,
			prevID,
		)

		if e != nil {
			utils.RecoverError()
			return e
		}

		//delete ans per user
		deleteDynStmt =
			`DELETE FROM public.answers ` +
				`WHERE form_answers_user_id = $1 ` +
				``
		_, e = db.Query(
			deleteDynStmt,
			prevID,
		)

		if e != nil {
			utils.RecoverError()
			return e
		}

		//delete ans per user
		deleteDynStmt =
			`DELETE FROM public.question_short_answer ` +
				`WHERE form_answers_user_id = $1 ` +
				``
		_, e = db.Query(
			deleteDynStmt,
			prevID,
		)

		if e != nil {
			utils.RecoverError()
			return e
		}
		//delete form_answers_user
		deleteDynStmt =
			`DELETE FROM public.form_answers_user ` +
				`WHERE id = $1 ` +
				``
		_, e = db.Query(
			deleteDynStmt,
			prevID,
		)

		if e != nil {
			utils.RecoverError()
			return e
		}
	}
	var relationalID int
	//save answers by user
	insertDynStmt :=
		`INSERT INTO public.form_answers_user ` +
			`(student_id, form_id,date) ` +
			`VALUES($1, $2,$3) ` +
			`RETURNING id;`
	rows, e := db.Query(
		insertDynStmt,
		answers.StudentId,
		answers.FormId,
		time.Now().UTC().Format("2006-01-02 15:04:05"),
	)

	if e != nil {
		utils.RecoverError()
		return e
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&relationalID,
		)
		if err != nil {
			return err
		}
	}

	//answers
	for _, answer := range answers.FormResponses {
		switch answer.QuestionType {
		case "RADIO":
			insertDynStmt =
				`INSERT INTO public.answers ` +
					`(form_answers_user_id,question_id,answers_selection_id) ` +
					`VALUES($1, $2, $3)`
			_, e := db.Exec(
				insertDynStmt,
				relationalID,
				answer.QuestionId,
				answer.AnswersSelectionId)

			if e != nil {
				utils.RecoverError()
				return e
			}
			continue
		case "LINEAR":
			insertDynStmt =
				`INSERT INTO public.answers ` +
					`(form_answers_user_id,question_id,answers_option_id) ` +
					`VALUES($1, $2, $3)`
			_, e := db.Exec(
				insertDynStmt,
				relationalID,
				answer.QuestionId,
				answer.AnswersItemId,
			)

			if e != nil {
				utils.RecoverError()
				return e
			}
			continue
		case "LINEAR_OPTION":
			insertDynStmt :=
				`INSERT INTO public.answers ` +
					`(form_answers_user_id,question_id,answers_option_id) ` +
					`VALUES($1, $2, $3)`
			_, e := db.Exec(
				insertDynStmt,
				relationalID,
				answer.QuestionId,
				answer.AnswersOptionId,
			)

			if e != nil {
				utils.RecoverError()
				return e
			}
			continue
		case "SHORT_ANSWER":
			var shortAnsID int
			ShortInsertDynStmt :=
				`INSERT INTO public.question_short_answer ` +
					`(form_answers_user_id,question_id,answer) ` +
					`VALUES($1, $2,$3) ` +
					`RETURNING id;`
			shortRows, e := db.Query(
				ShortInsertDynStmt,
				relationalID,
				answer.QuestionId,
				answer.AnswersShortQuestion,
			)

			if e != nil {
				utils.RecoverError()
				return e
			}
			defer shortRows.Close()
			for shortRows.Next() {
				err = shortRows.Scan(
					&shortAnsID,
				)
				if err != nil {
					return err
				}
			}
			insertDynStmt =
				`INSERT INTO public.answers ` +
					`(form_answers_user_id,question_id,answer_short_id) ` +
					`VALUES($1, $2,$3)`
			_, e = db.Exec(
				insertDynStmt,
				relationalID,
				answer.QuestionId,
				shortAnsID,
			)

			if e != nil {
				utils.RecoverError()
				return e
			}
			continue
		}

	}

	//time per sections
	for _, timeSection := range answers.SectionTime {
		insertDynStmt :=
			`INSERT INTO public.time_per_section ` +
				`("user",section_id,time_in_seconds) ` +
				`VALUES($1, $2, $3)`
		_, e := db.Exec(
			insertDynStmt,
			relationalID,
			timeSection.SectionID,
			timeSection.SectionTime,
		)

		if e != nil {
			utils.RecoverError()
			return e
		}
	}

	return nil
}

func GetAnswers(answers models.FormResponse, db *sql.DB, userID, formID int) (FormResponse models.FormResponse, err error) {
	var qli map[int]string

	FormResponse.FormId = formID
	FormResponse.StudentId = userID
	var idFormResponse int

	//obtencion de tipos de respuesta de acuerdo etc.
	qli = make(map[int]string)
	rowsItem, err := db.Query(`SELECT "order",item ` +
		`FROM public.question_linear_item `)

	if err != nil {
		panic(err)
	}

	defer rowsItem.Close()
	for rowsItem.Next() {
		var (
			order int
			item  string
		)
		err = rowsItem.Scan(
			&order,
			&item,
		)
		if err == nil {
			qli[order] = item
		}

	}
	// obtencion de formularios respondidos por el usuario
	var FormResponses []int
	rows, err := db.Query(`SELECT `+
		`id,date `+
		`FROM public.form_answers_user as fau `+
		`WHERE fau.student_id = $1 and fau.form_id = $2`, userID, formID)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(
			&idFormResponse,
			&FormResponse.Date,
		)
		if err == nil {
			FormResponses = append(FormResponses, idFormResponse)
		}

	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	//obtencion de respuestas por formulario
	for _, formResponses := range FormResponses {
		var answers models.FormResponses
		answers.IsCorrect = nil
		rows, err = db.Query(`SELECT `+
			`COALESCE(q.description,'sin descripcion'),`+
			`COALESCE(ans.answers_option_id,0),`+
			`COALESCE(ans.answers_selection_id,0), `+
			`COALESCE(ans.answers_item_id,0), `+
			`COALESCE(ans.answer_short_id,0), `+
			`ans.assigne_score, `+
			`ans.question_id, `+
			`sq.section_id, `+
			`s.title, `+
			`s.score_for_each_question, `+
			`q.is_open_question, `+
			`qt.type, `+
			`q.has_score, `+
			`COALESCE(q.title,''), `+
			`COALESCE(q.question_description,''),  `+
			`COALESCE(q.image_url,'') `+
			`FROM public.answers ans `+
			`INNER JOIN public.question q ON q.id = ans.question_id `+
			`INNER JOIN public.question_type qt ON qt.id = q.type_id `+
			`INNER JOIN public.section_questions sq ON q.id = sq.question_id `+
			`INNER JOIN public.section s ON s.id = sq.section_id `+
			`WHERE ans.form_answers_user_id = $1 ORDER BY question_id ASC `, formResponses)
		if err != nil {
			panic(err)
		}

		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(
				&answers.Question,
				&answers.AnswersOptionId,
				&answers.AnswersSelectionId,
				&answers.AnswersItemId,
				&answers.AnswersShortQuestionId,
				&answers.AssigneScore,
				&answers.QuestionId,
				&answers.SectionId,
				&answers.SectionTitle,
				&answers.ScoreForEachQuestion,
				&answers.IsOpenQuestion,
				&answers.QuestionType,
				&answers.HasScore,
				&answers.Title,
				&answers.QuestionDescription,
				&answers.ImgURL,
			)
			if err == nil {
				if *answers.ImgURL != "" {
					answers.QuestionHasImage = true
				}
				if *answers.Title == "" {
					answers.Title = nil
				}
				if *answers.QuestionDescription == "" {
					answers.QuestionDescription = nil
				}
				if answers.AnswersSelectionId != 0 {
					rowsPerAns, err := db.Query(`SELECT `+
						`qso.option,is_correct `+
						`FROM public.question_selection_option qso `+
						`INNER JOIN public.answers a ON a.answers_selection_id = qso.id `+
						`WHERE a.answers_selection_id = $1 `, answers.AnswersSelectionId)
					if err != nil {
						panic(err)
					}

					defer rowsPerAns.Close()
					for rowsPerAns.Next() {
						_ = rowsPerAns.Scan(
							&answers.Answer,
							&answers.IsCorrect,
						)
					}
					if err = rowsPerAns.Err(); err != nil {
						panic(err)
					}
				}
				if answers.AnswersOptionId != 0 {
					rowsPerAns, err := db.Query(`SELECT `+
						`COALESCE(q.description,''),qso.option, qso.is_correct,qso.image_url `+
						`FROM public.question_linear_option qso `+
						`INNER JOIN public.answers a ON a.answers_option_id = qso.id `+
						`INNER JOIN public.question_linear q ON q.question_id = qso.question_id `+
						`WHERE a.answers_option_id = $1 `, answers.AnswersOptionId)
					if err != nil {
						panic(err)
					}

					defer rowsPerAns.Close()
					for rowsPerAns.Next() {
						_ = rowsPerAns.Scan(
							&answers.Question,
							&answers.Answer,
							&answers.IsCorrect,
							&answers.ImgURL,
						)
					}
					if err = rowsPerAns.Err(); err != nil {
						panic(err)
					}
					answerInt, err := strconv.Atoi(answers.Answer)
					if err == nil {
						answers.Answer = qli[answerInt]
					}
				}
				if answers.AnswersItemId != 0 {
					rowsPerAns, err := db.Query(`SELECT `+
						`qso.option `+
						`FROM public.question_linear_item qso `+
						`INNER JOIN public.answers a ON a.answers_item_id = qso.id `+
						`WHERE a.answers_item_id = $1 `, answers.AnswersItemId)
					if err != nil {
						panic(err)
					}

					defer rowsPerAns.Close()
					for rowsPerAns.Next() {
						_ = rowsPerAns.Scan(
							&answers.Answer,
						)
					}
					if err = rowsPerAns.Err(); err != nil {
						panic(err)
					}
				}
				if answers.AnswersShortQuestionId != 0 {
					answers.IsCorrect = nil
					rowsPerAns, err := db.Query(`SELECT `+
						`qsa.answer `+
						`FROM public.question_short_answer qsa `+
						`WHERE id = $1 `, answers.AnswersShortQuestionId)
					if err != nil {
						panic(err)
					}

					defer rowsPerAns.Close()
					for rowsPerAns.Next() {
						_ = rowsPerAns.Scan(
							&answers.Answer,
						)
					}
					if err = rowsPerAns.Err(); err != nil {
						panic(err)
					}
				}
				FormResponse.FormResponses = append(FormResponse.FormResponses, answers)
			}

		}
		if err = rows.Err(); err != nil {
			panic(err)
		}
	}

	//obtencion de tiempos por seccion
	rows, err = db.Query(`SELECT `+
		`section_id,`+
		`time_in_seconds `+
		`FROM public.time_per_section as tps `+
		`WHERE tps."user" = $1`, idFormResponse)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var SectionTime models.SectionTime
		err = rows.Scan(
			&SectionTime.SectionID,
			&SectionTime.SectionTime,
		)
		if err == nil {
			FormResponse.SectionTime = append(FormResponse.SectionTime, SectionTime)
		}

	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}

func GetFormAns(answers models.FormResponse, db *sql.DB, userID int) (Forms models.Forms, err error) {
	rows, err := db.Query(`SELECT `+
		`f.id,f.title,fau.date `+
		`FROM public.form_answers_user as fau `+
		`INNER JOIN public.form f ON fau.form_id = f .id `+
		`WHERE fau.student_id = $1 `, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var form models.Form

		err = rows.Scan(
			&form.FormId,
			&form.FormTitle,
			&form.FormDate,
		)
		if err == nil {
			Forms.FormId = append(Forms.FormId, form)
		}

	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}

func GetAllFormAns(answers models.FormResponse, db *sql.DB) (Forms models.Forms, err error) {
	rows, err := db.Query(`SELECT ` +
		`s.id,s.full_name,c.name,f.id,f.title,fau.date ` +
		`FROM public.form_answers_user as fau ` +
		`INNER JOIN public.student s ON s.id = fau.student_id ` +
		`INNER JOIN public.career c ON c.id = s.career_id ` +
		`INNER JOIN public.form f ON fau.form_id = f .id `)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var form models.Form

		err = rows.Scan(
			&form.StudentId,
			&form.StudentName,
			&form.CarrerName,
			&form.FormId,
			&form.FormTitle,
			&form.FormDate,
		)
		if err == nil {
			Forms.FormId = append(Forms.FormId, form)
		}

	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}

func GetAllFormAnsByFormId(answers models.FormResponse, db *sql.DB, formId int) (Forms models.Forms, err error) {
	rows, err := db.Query(`SELECT `+
		`s.id,s.full_name,c.name,f.id,f.title,fau.date `+
		`FROM public.form_answers_user as fau `+
		`INNER JOIN public.student s ON s.id = fau.student_id `+
		`INNER JOIN public.career c ON c.id = s.career_id `+
		`INNER JOIN public.form f ON fau.form_id = f .id `+
		`WHERE fau.form_id = $1 `, formId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var form models.Form

		err = rows.Scan(
			&form.StudentId,
			&form.StudentName,
			&form.CarrerName,
			&form.FormId,
			&form.FormTitle,
			&form.FormDate,
		)
		if err == nil {
			Forms.FormId = append(Forms.FormId, form)
		}

	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	return
}

func InsertAssigneScore(assigneScore models.AssigneScore, db *sql.DB) (err error) {
	// GET QUESTION_ANS ID
	var relationalID int
	rows, err := db.Query(`SELECT `+
		`id `+
		`FROM public.form_answers_user as fau `+
		`WHERE fau.form_id = $1 and student_id = $2`, assigneScore.FormId, assigneScore.StudentId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(
			&relationalID,
		)
	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	// INSERT SCORE
	insertDynStmt :=
		`UPDATE public.answers ` +
			`SET assigne_score = $1 ` +
			`WHERE question_id = $2 AND form_answers_user_id = $3  `
	rows, e := db.Query(
		insertDynStmt,
		assigneScore.AssigneScore,
		assigneScore.QuestionId,
		relationalID,
	)

	if e != nil {
		utils.RecoverError()
		return e
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(
			&relationalID,
		)
		if err != nil {
			return err
		}
	}
	return
}

func GetIndicatorsSql(db *sql.DB) (form models.Indicators, err error) {

	var CarrersIndicator models.CarrerIndicator
	rows, err := db.Query(`SELECT ` +
		`g.gender_name,f.id,f.title ` +
		`FROM public.form_answers_user as fau ` +
		`INNER JOIN public.student s ON s.id = fau.student_id ` +
		`INNER JOIN public.gender g ON g.id = s.gender_id ` +
		`INNER JOIN public.form f ON fau.form_id = f .id `)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {

		var genderType string

		err = rows.Scan(
			&genderType,
			&form.FormID,
			&form.FormName,
		)
		if err == nil {
			form.Gender.Total++
			if genderType == "male" {
				form.Gender.Men++
			} else {
				form.Gender.Women++
			}
		}

	}
	if err = rows.Err(); err != nil {
		panic(err)
	}

	rows, err = db.Query(`SELECT ` +
		`c.name,c.id ` +
		`FROM public.form_answers_user as fau ` +
		`INNER JOIN public.student s ON s.id = fau.student_id ` +
		`INNER JOIN public.career c ON c.id = s.career_id ` +
		`INNER JOIN public.form f ON fau.form_id = f .id `)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var CarrerIndicator models.Carrers

		err = rows.Scan(
			&CarrerIndicator.Career,
			&CarrerIndicator.CareerID,
		)
		if err == nil {
			CarrersIndicator.Total++
			CarrersIndicator.Carrers = append(CarrersIndicator.Carrers, CarrerIndicator)
		}

	}
	if err = rows.Err(); err != nil {
		panic(err)
	}
	form.Carrer = CarrersIndicator

	return
}
