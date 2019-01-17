package country

import (
	"errors"

	"github.com/7phs/coding-challenge-iban/helper"
)

var (
	ErrCountryLength   = errors.New("length invalid")
	ErrCountryTemplate = errors.New("not match country template")
	ErrCountryKk       = errors.New("kk invalid")
)

type Format struct {
	Country  string       `yaml:"country"`
	Len      int          `yaml:"len"`
	Template IbanTemplate `yaml:"template"`
	Kk       string       `yaml:"kk"`
}

func (o *Format) Validate(kk, text, suffix string) error {
	var errList helper.ErrList

	if len(text) != o.Len {
		errList.Add(ErrCountryLength)
	}

	if o.Kk != "" && o.Kk != kk {
		errList.Add(ErrCountryKk)
	}

	if !o.Template.Match([]byte(suffix)) {
		errList.Add(ErrCountryTemplate)
	}

	return errList.Result()
}
