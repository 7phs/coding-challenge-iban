package validator

import "github.com/7phs/coding-challenge-iban/model/records"

type Status string

const (
	Valid   Status = "valid"
	Invalid Status = "invalid"
)

type Validator interface {
	Validate(*records.Iban) error
}
