package consumer

import (
	"net/http"
	"service-xyz/config"
	"service-xyz/internal/usecase/consumer"
	"service-xyz/pkg/logger"

	"github.com/gorilla/mux"
)

type ConsumerRoutes struct {
	l   *logger.Logger
	cfg *config.Config
	cu  consumer.IConsumerUsecase
}

func NewConsumerRoutes(r *mux.Router, l *logger.Logger, cfg *config.Config, cu consumer.IConsumerUsecase) {
	c := &ConsumerRoutes{l, cfg, cu}

	group := r.PathPrefix("/v1/consumer").Subrouter()
	group.HandleFunc("/create", c.CreateConsumer).Methods(http.MethodPost)
	group.HandleFunc("/get/{consumer_id}", c.GetConsumerById).Methods(http.MethodGet)
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("OK Service Running.."))
	}).Methods(http.MethodGet)
}
