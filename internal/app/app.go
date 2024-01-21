package app

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"service-xyz/config"
	v1 "service-xyz/internal/controller/http/v1"
	mysql_repository "service-xyz/internal/repository/mysql"
	"service-xyz/internal/usecase/consumer"
	"service-xyz/pkg/httpserver"
	"service-xyz/pkg/logger"
	"service-xyz/pkg/mysql"
	"syscall"

	"github.com/gorilla/mux"
)

func Run(cfg *config.Config) {
	fmt.Println("Running Service-XYZ")

	var err error
	l := logger.New(cfg)

	// Mysql
	db := mysql.New(cfg.MYSQL.MysqlDriverName, cfg, l)
	defer db.Close()

	// Repository
	consumerRepository := mysql_repository.NewConsumerMysqlRepository(l, cfg, db)

	// Usecase
	consumerUsecase := consumer.NewConsumerUsecase(l, cfg, consumerRepository)

	// HTTP Server
	handler := mux.NewRouter()
	v1.NewRouter(handler, l, cfg, consumerUsecase)
	httpServer := httpserver.New(handler, cfg, httpserver.Port(cfg.HTTPServer.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Println(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Println(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
