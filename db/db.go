package db

import (
	"fmt"
	"github.com/7phs/coding-challenge-iban/config"
	"github.com/7phs/coding-challenge-iban/db/file"
	log "github.com/sirupsen/logrus"
)

var (
	CountriesFmt CountriesFormat
)

func Init(config *config.Config) (err error) {
	log.Debug("db: countries format init")

	CountriesFmt, err = file.NewCountriesFormat(config.DbPath())
	if err!=nil {
		err = fmt.Errorf("db: countries format, %v", err)
	}

	return err
}

func Shutdown() {
}