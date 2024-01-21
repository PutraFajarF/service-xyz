package consumer

import (
	"encoding/json"
	"net/http"
	"service-xyz/commons"
	http_resp "service-xyz/internal/controller/response"
	"service-xyz/internal/entity"
	"service-xyz/pkg/logger"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *ConsumerRoutes) CreateConsumer(w http.ResponseWriter, r *http.Request) {
	var payload entity.ConsumerInfoRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		http_resp.HttpErrorResponse(w, false, http.StatusBadRequest, "400", err.Error())
		return
	}

	err := c.cu.CreateConsumer(&payload)
	if err != nil {
		http_resp.HttpErrorResponse(w, false, http.StatusInternalServerError, "500", commons.ErrInternalServer.Error())
		return
	}

	c.l.CreateLog(&logger.Log{
		Event:      commons.HTTP_CONSUMER + "|createConsumer",
		Method:     "POST",
		StatusCode: http.StatusOK,
		Request:    payload,
		Response:   "Success",
		Message:    "Success create consumer",
	}, logger.LVL_INFO)

	http_resp.HttpSuccessResponse(w, true, http.StatusOK, "200", commons.HTTP_CONSUMER+" Create Success", nil)
}

func (c *ConsumerRoutes) GetConsumerById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	consumerId := param["consumer_id"]
	intConsumerId, err := strconv.Atoi(consumerId)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.HTTP_CONSUMER + "|get Consumer",
			Method:     "GET",
			StatusCode: http.StatusBadRequest,
			Request:    "consumer_id: " + consumerId,
			Response:   err,
			Message:    "Fail get consumer",
		}, logger.LVL_ERROR)
		http_resp.HttpErrorResponse(w, false, http.StatusBadRequest, "400", err.Error())
		return
	}

	res, err := c.cu.GetDataConsumerById(intConsumerId)
	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.HTTP_CONSUMER + "|get Consumer",
			Method:     "GET",
			StatusCode: http.StatusInternalServerError,
			Request:    "consumer_id: " + consumerId,
			Response:   err,
			Message:    "Fail get consumer",
		}, logger.LVL_ERROR)
		http_resp.HttpErrorResponse(w, false, http.StatusInternalServerError, "500", err.Error())
		return
	}

	c.l.CreateLog(&logger.Log{
		Event:      commons.HTTP_CONSUMER + "|getConsumer",
		Method:     "GET",
		StatusCode: http.StatusOK,
		Request:    consumerId,
		Response:   "Success",
		Message:    "Success get consumer",
	}, logger.LVL_INFO)

	http_resp.HttpSuccessResponse(w, true, http.StatusOK, "200", commons.HTTP_CONSUMER+" Get Success", res)
}
