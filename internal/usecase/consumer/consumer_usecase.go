package consumer

import (
	"encoding/json"
	"net/http"
	"service-xyz/commons"
	"service-xyz/config"
	"service-xyz/internal/entity"
	"service-xyz/internal/repository/mysql"
	"service-xyz/pkg/logger"
	"sync"
	"time"
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
	var wg sync.WaitGroup
	var mu sync.Mutex
	var errCh error

	request := map[string]interface{}{"request": data}
	jsonReq, _ := json.Marshal(request)

	Consumer := entity.ConsumerInfo{
		NIK:          data.NIK,
		Email:        data.Email,
		Gender:       data.Gender,
		FullName:     data.FullName,
		LegalName:    data.LegalName,
		TempatLahir:  data.TempatLahir,
		TanggalLahir: data.TanggalLahir,
		Gaji:         data.Gaji,
		FotoKTP:      data.FotoKTP,
		FotoSelfie:   data.FotoSelfie,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Gunakan Mutex sebelum memanggil InsertConsumer untuk hindari race condition
	mu.Lock()
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Memanggil InsertConsumer di dalam goroutine
		err := c.cr.InsertConsumer(&Consumer)
		if err != nil {
			errCh = err

			defer c.l.CreateLog(&logger.Log{
				Event:      commons.USECASE_CONSUMER + "|CREATE",
				Method:     "POST",
				StatusCode: http.StatusInternalServerError,
				Request:    string(jsonReq),
				Query:      "",
				Response:   err,
				Message:    commons.ErrUsecaseConsumer,
			}, logger.LVL_ERROR)
		}
	}()
	mu.Unlock()

	// Wait goroutine selesai
	wg.Wait()

	if errCh != nil {
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
			Event:      commons.USECASE_CONSUMER + "|GET",
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
