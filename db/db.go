package db

import (
	"fmt"

	"github.com/7phs/coding-challenge-iban/config"
	"github.com/7phs/coding-challenge-iban/db/file"
	"github.com/7phs/coding-challenge-iban/model/validator"
	log "github.com/sirupsen/logrus"
)

var (
	CountriesFmt CountriesFormat
)

type DB interface {
	Shutdown()
}

type CountriesFormat interface {
	DB

	validator.CountriesFormatDB
}

func Init(config *config.Config) (err error) {
	log.Debug("db: countries format init")

	CountriesFmt, err = file.NewCountriesFormat(config.DbPath())
	if err != nil {
		err = fmt.Errorf("db: countries format, %v", err)
	}

	return err
}

func Shutdown() {
}
