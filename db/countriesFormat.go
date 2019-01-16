package db

import (
	"github.com/7phs/coding-challenge-iban/model/validator"
)

type DB interface {
	Shutdown()
}

type CountriesFormat interface {
	DB

	validator.CountriesFormatDB
}
