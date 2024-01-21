package consumer

import (
	"fmt"
	"service-xyz/config"
	"service-xyz/internal/entity"
	"service-xyz/internal/usecase/consumer/mocks"
	"service-xyz/pkg/logger"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CreateConsumer_ShouldSuccess(t *testing.T) {
	config, _ := config.NewConfig()
	logger := logger.New(config)
	mockConsumerRepo := mocks.IConsumerMysqlRepository{}

	consumerUsecase := NewConsumerUsecase(logger, config, &mockConsumerRepo)
	mockConsumerRepo.On("InsertConsumer", mock.Anything).Return(nil)

	data := &entity.ConsumerInfoRequest{}
	err := consumerUsecase.CreateConsumer(data)

	assert.Equal(t, err, nil, "Success Create Consumer")
}

func Test_GetDataConsumerById_ShouldSuccess(t *testing.T) {
	config, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}
	logger := logger.New(config)
	mockConsumerRepo := mocks.IConsumerMysqlRepository{}

	consumerUsecase := NewConsumerUsecase(logger, config, &mockConsumerRepo)
	mockConsumerRepo.On("GetConsumerById", mock.Anything).Return(&entity.ConsumerInfo{}, nil)

	data := 1
	res, _ := consumerUsecase.GetDataConsumerById(data)

	assert.Equal(t, res, &entity.ConsumerInfo{}, "Fail Create Consumer")
}
