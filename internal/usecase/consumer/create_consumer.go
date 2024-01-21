package consumer

import (
	"encoding/json"
	"net/http"
	"service-xyz/commons"
	"service-xyz/config"
	"service-xyz/internal/entity"
	"service-xyz/internal/repository/mysql"
	"service-xyz/pkg/logger"
)

type IConsumerUsecase interface {
	CreateConsumer(data *entity.ConsumerInfoRequest) error
	GetDataConsumerById(id int) (*entity.ConsumerInfo, error)
}

type ConsumerUseCase struct {
	l   *logger.Logger
	cfg *config.Config
	cr  mysql.IConsumerMysqlRepository
}

func NewConsumerUsecase(l *logger.Logger, cfg *config.Config, cr mysql.IConsumerMysqlRepository) *ConsumerUseCase {
	return &ConsumerUseCase{l, cfg, cr}
}

func (c *ConsumerUseCase) CreateConsumer(data *entity.ConsumerInfoRequest) error {
	request := map[string]interface{}{"request": data}
	jsonReq, _ := json.Marshal(request)

	err := c.cr.InsertConsumer(data)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.USECASE_CONSUMER + "|CREATE",
			Method:     "POST",
			StatusCode: http.StatusInternalServerError,
			Request:    string(jsonReq),
			Query:      "",
			Response:   err,
			Message:    commons.ErrUsecaseConsumer,
		}, logger.LVL_ERROR)
		return commons.ErrUsecaseConsumer
	}

	return nil
}

func (c *ConsumerUseCase) GetDataConsumerById(id int) (*entity.ConsumerInfo, error) {
	request := map[string]interface{}{"customer_id": id}
	jsonReq, _ := json.Marshal(request)

	result, err := c.cr.GetConsumerById(id)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.USECASE_CONSUMER + "|CREATE",
			Method:     "GET",
			StatusCode: http.StatusInternalServerError,
			Request:    string(jsonReq),
			Query:      "",
			Response:   err,
			Message:    commons.ErrUsecaseConsumer,
		}, logger.LVL_ERROR)
		return nil, commons.ErrUsecaseConsumer
	}

	return result, nil
}
