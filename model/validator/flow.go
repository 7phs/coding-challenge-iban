package validator

import "github.com/7phs/coding-challenge-iban/model/records"

type Flow struct {
	validators []Validator
}

func NewFlow() *Flow {
	return &Flow{}
}

func (o *Flow) Then(validator Validator) *Flow {
	o.validators = append(o.validators, validator)

	return o
}

func (o *Flow) Validate(rec *records.Iban) error {
	for _, validator := range o.validators {
		if err := validator.Validate(rec); err != nil {
			return err
		}
	}

	return nil
}
