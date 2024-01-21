package v1

import (
	"service-xyz/config"
	consumer_http "service-xyz/internal/controller/http/v1/consumer"
	"service-xyz/internal/usecase/consumer"
	"service-xyz/pkg/logger"

	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, l *logger.Logger, cfg *config.Config, cu consumer.IConsumerUsecase) {
	{
		consumer_http.NewConsumerRoutes(r, l, cfg, cu)
	}
}
