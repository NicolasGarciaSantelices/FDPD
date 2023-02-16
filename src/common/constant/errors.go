package constat_errors

import "errors"

// database errors
var ErrorEmptyData = errors.New("no data return from database")
var ErrorConnect = errors.New("failed to connect to database")
var ErrorInRequestBody = errors.New("the body sent cannot be read correctly, multiple fields with error")
var ErrorInDateParams = errors.New("the dates sent are incorrect, please check the format")
var ErrorFromOrDate = errors.New("the date entered for 'from' has to be less than 'to'")
var ErrorFrom = errors.New("format error in 'from'")
var ErrorTo = errors.New("format error in 'to")
var ErrorInLicensePlate = errors.New("license plate field is empty")
var ErrorInBattery = errors.New("baterry field is empty")
var ErrorChart = errors.New("invalid chart, available chart types are chart speed, chart soc, chart throttle")
var ErrorEmptyEndpointData = errors.New("no data")
var ErrorInFileJSON = errors.New("format error in file")
var ErrorInUpdates = errors.New("error in ins")
var ErrorInQuery = errors.New("error in query params")
