package validator

import (
	"os"
	"regexp"

	"github.com/7phs/coding-challenge-iban/model/records"
)

var (
	matcher = regexp.MustCompile(`^([a-zA-Z0-9]+\s*)+$`)
)

type Raw struct{}

func NewRaw() *Raw {
	return &Raw{}
}

func (o *Raw) Validate(rec *records.Iban) error {
	if !matcher.Match([]byte(rec.Raw())) {
		return os.ErrInvalid
	}

	return nil
}
