package validator

import (
	"testing"

	"github.com/7phs/coding-challenge-iban/db/file"
	"github.com/7phs/coding-challenge-iban/model/records"
	"github.com/stretchr/testify/assert"
)

func TestCountry(t *testing.T) {
	testSuites := []*struct {
		in  string
		exp bool
	}{
		{
			in:  "BR97 0036 0305 0000 1000 9795 493P 1",
			exp: true,
		},
		{
			in:  "GB82WEST12345698765432",
			exp: false,
		},
		{
			in:  "BR97AA82WEST12345698765432",
			exp: false,
		},
	}

	countriesDB, err := file.NewCountriesFormat("../../test-data/countries-test.yaml")
	if !assert.NoError(t, err, "failed to read a db") {
		return
	}

	validator := NewCountry(countriesDB)

	for i, test := range testSuites {
		err := validator.Validate(records.NewIban(test.in))
		if !test.exp {
			assert.Error(t, err, "%d: failed to catch error", i)
		} else {
			assert.NoError(t, err, "%d: wrong validate and got %v", i, err)
		}
	}

}
