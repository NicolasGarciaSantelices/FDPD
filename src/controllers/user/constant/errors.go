package constant

import "errors"

// database errors
var ErrorInAuth = errors.New("user_does_not_belong_to_ucn")
var ErrorUserExist = errors.New("user_already_exists")
var ErrorUserUpdate = errors.New("error_when_try_to_update_user_info")
var InsertAnsError = "unable_to_insert_answers_check_json_body"
var InsertAnsErrorStatus = "Error"
var InsertAndsSuccesStatus = "successfully"
var InsertAns = "insert_answers_and_time_per_section_succesfully"
var SuccesStatus = "successfully"
var GetAnsSucces = "get_answers_and_time_per_section_succesfully"
var GetAnsFormSucces = "get_answers_forms_succesfully"
var GetIndicatorsSucces = "get_indicator_succesfully"
