package model

import (
	"github.com/7phs/coding-challenge-iban/model/validator"
)

var (
	valid validator.Validator
)

type Dependencies struct {
	CountriesFormat validator.CountriesFormatDB
}

func Init(deps *Dependencies) {
	valid = validator.NewFlow().
		Then(validator.NewRaw()).
		Then(validator.NewCountry(deps.CountriesFormat)).
		Then(validator.NewCheckSum())
}

func Validator() validator.Validator {
	return valid
}
