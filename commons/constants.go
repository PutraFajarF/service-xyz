package commons

import "errors"

var (
	// HTTP
	HTTP_CONSUMER = "HTTP|CONSUMER"

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
