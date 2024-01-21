package commons

import "errors"

var (
	// HTTP
	HTTP_CONSUMER           = "HTTP|CONSUMER"
	SUCCESS_CREATE_CONSUMER = "Success Create Consumer"
	FAIL_CREATE_CONSUMER    = "Fail Create Consumer"
	SUCCESS_GET_CONSUMER    = "Success Get Consumer"
	FAIL_GET_CONSUMER       = "Fail Get Consumer"

	// USECASE CONSUMER
	USECASE_CONSUMER = "USECASE|CONSUMER"
	// REPOSITORY MYSQL
	REPOSITORY_MYSQL_CONSUMER = "REPOSITORY|MYSQL|CONSUMER"

	// REPOSITORY MYSQL ERROR
	ErrQuery        = errors.New("error - execute query")
	ErrPrepareQuery = errors.New("error - preparing query statement")
	ErrRowScan      = errors.New("error - scanning rows repository")

	// USECASE CONSUMER ERROR
	ErrUsecaseConsumer = errors.New("error - usecase consumer")

	// COMMON ERROR
	ErrInvalidPayload = errors.New("error - invalid request payload")
	ErrInternalServer = errors.New("error - internal server error")
)
