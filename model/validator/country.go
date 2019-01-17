package validator

import (
	"github.com/7phs/coding-challenge-iban/model/records"
	"github.com/7phs/coding-challenge-iban/model/validator/country"
)

type CountriesFormatDB interface {
	Get(countryCode string) (*country.Format, error)
}

type Country struct {
	db CountriesFormatDB
}

func NewCountry(db CountriesFormatDB) *Country {
	return &Country{
		db: db,
	}
}

func (o *Country) Validate(rec *records.Iban) error {
	fmt, err := o.db.Get(rec.CountryCode())
	if err != nil {
		return err
	}

	return fmt.Validate(rec.Kk(), rec.Text(), rec.Suffix())
}
