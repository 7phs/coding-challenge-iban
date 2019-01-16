package file

import (
	"bufio"
	"github.com/7phs/coding-challenge-iban/model/validator/country"
	"gopkg.in/yaml.v2"
	"os"
)

type CountriesFormat struct {
	databasePath string

	records map[string]*country.Format
}

func NewCountriesFormat(databasePath string) (*CountriesFormat, error) {
	return (&CountriesFormat{
		databasePath: databasePath,
	}).Init()
}

func (o *CountriesFormat) Init() (*CountriesFormat, error) {
	file, err := os.Open(o.databasePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(bufio.NewReader(file))

	err = decoder.Decode(&o.records)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (o *CountriesFormat) Get(countryCode string) (*country.Format, error) {
	v, ok := o.records[countryCode]
	if !ok {
		return nil, os.ErrNotExist
	}

	return v, nil
}

func (o *CountriesFormat) Shutdown() {}
