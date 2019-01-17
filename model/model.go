package model

import (
	"github.com/7phs/coding-challenge-iban/model/validator"
)

var (
	Default Models
)

type Models interface {
	Validator() validator.Validator
}

type Dependencies struct {
	CountriesFormat validator.CountriesFormatDB
}

type DefaultState struct {
	Valid validator.Validator
}

func (o *DefaultState) Validator() validator.Validator {
	return o.Valid
}

func DefaultValidator(deps *Dependencies) *validator.Flow {
	return validator.NewFlow().
		Then(validator.NewRaw()).
		Then(validator.NewCountry(deps.CountriesFormat)).
		Then(validator.NewCheckSum())
}

func Init(deps *Dependencies) {
	Default = &DefaultState{
		Valid: DefaultValidator(deps),
	}
}
