package constant

import "errors"

// database errors
var ErrorInAuth = errors.New("user_does_not_belong_to_ucn")
var ErrorUserExist = errors.New("user_already_exists")
var ErrorUserUpdate = errors.New("error_when_try_to_update_user_info")
