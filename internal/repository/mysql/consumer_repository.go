package mysql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"service-xyz/commons"
	"service-xyz/config"
	"service-xyz/internal/entity"
	"service-xyz/pkg/logger"
)

type IConsumerMysqlRepository interface {
	InsertConsumer(data *entity.ConsumerInfoRequest) error
	GetConsumerById(id int) (*entity.ConsumerInfo, error)
	// UpdateConsumerById(id int) (*entity.ConsumerInfo, error)
}

type ConsumerMysqlRepo struct {
	l   *logger.Logger
	cfg *config.Config
	db  *sql.DB
}

func NewConsumerMysqlRepository(l *logger.Logger, cfg *config.Config, db *sql.DB) *ConsumerMysqlRepo {
	return &ConsumerMysqlRepo{l, cfg, db}
}

func (c *ConsumerMysqlRepo) InsertConsumer(data *entity.ConsumerInfoRequest) error {
	jsonReq, _ := json.Marshal(data)

	trx, err := c.db.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO consumer_info (
		nik,
		email,
		gender,
		full_name,
		legal_name,
		tempat_lahir,
		tanggal_lahir,
		gaji,
		foto_ktp,
		foto_selfie,
		created_at,
		updated_at,
	)
	values (AES_ENCRYPT(?,'%[1]s'), AES_ENCRYPT(?,'%[1]s'), ?, ?, ?, ?, ?, ?, AES_ENCRYPT(?,'%[1]s'), ?, ?, ?)`

	query = fmt.Sprintf(query, c.cfg.Cipher.CipherMysql)
	stmt, err := trx.Prepare(query)

	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.REPOSITORY_MYSQL_CONSUMER + "|prepareQuery",
			Method:     "",
			StatusCode: http.StatusInternalServerError,
			Request:    string(jsonReq),
			Query:      query,
			Response:   err,
			Message:    commons.ErrQuery,
		}, logger.LVL_ERROR)
		return commons.ErrPrepareQuery
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		data.NIK,
		data.Email,
		data.Gender,
		data.FullName,
		data.LegalName,
		data.TempatLahir,
		data.TanggalLahir,
		data.Gaji,
		data.FotoKTP,
		data.FotoSelfie,
		data.CreatedAt,
		data.UpdatedAt,
	)

	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.REPOSITORY_MYSQL_CONSUMER + "|execQuery",
			Method:     "",
			StatusCode: http.StatusInternalServerError,
			Request:    string(jsonReq),
			Query:      query,
			Response:   err,
			Message:    commons.ErrQuery,
		}, logger.LVL_ERROR)
		return commons.ErrPrepareQuery
	}

	if row, err := result.RowsAffected(); err != nil && row <= 0 {
		_ = trx.Rollback()
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.REPOSITORY_MYSQL_CONSUMER + "|rollbackQuery",
			Method:     "",
			StatusCode: http.StatusInternalServerError,
			Request:    string(jsonReq),
			Query:      query,
			Response:   err,
			Message:    commons.ErrQuery,
		}, logger.LVL_ERROR)
		return commons.ErrPrepareQuery
	}

	trx.Commit()

	return nil
}

func (c *ConsumerMysqlRepo) GetConsumerById(id int) (*entity.ConsumerInfo, error) {
	var res entity.ConsumerInfo

	request := map[string]int{"consumerId": id}
	jsonReq, _ := json.Marshal(request)

	query := "SELECT '%[2]d', AES_DECRYPT(nik,'%[1]s') nik, AES_DECRYPT(email,'%[1]s') email, gender, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, AES_DECRYPT(foto_ktp,'%[1]s') foto_ktp, foto_selfie, created_at, updated_at"
	query = fmt.Sprintf(query, c.cfg.Cipher.CipherMysql, id)
	prep, err := c.db.Prepare(query)

	if err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.REPOSITORY_MYSQL_CONSUMER + "|prepareQuery",
			Method:     "",
			StatusCode: http.StatusInternalServerError,
			Request:    string(jsonReq),
			Query:      query,
			Response:   err,
			Message:    commons.ErrQuery,
		}, logger.LVL_ERROR)
		return nil, commons.ErrPrepareQuery
	}

	defer prep.Close()

	row := c.db.QueryRow(query, id)

	if err := row.Scan(
		&res.Id,
		&res.NIK,
		&res.Email,
		&res.Gender,
		&res.FullName,
		&res.LegalName,
		&res.TempatLahir,
		&res.TanggalLahir,
		&res.Gaji,
		&res.FotoKTP,
		&res.FotoSelfie,
		&res.CreatedAt,
		&res.UpdatedAt,
	); err != nil {
		defer c.l.CreateLog(&logger.Log{
			Event:      commons.REPOSITORY_MYSQL_CONSUMER + "|rowScan",
			Method:     "",
			StatusCode: http.StatusInternalServerError,
			Request:    string(jsonReq),
			Query:      query,
			Response:   err,
			Message:    commons.ErrQuery,
		}, logger.LVL_ERROR)
		return nil, commons.ErrRowScan
	}

	return &res, nil
}
