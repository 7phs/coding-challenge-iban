package file

import (
	"testing"

	"github.com/7phs/coding-challenge-iban/model/validator/country"
	"github.com/stretchr/testify/assert"
)

func testParseTemplate(t *testing.T, template string) (rec country.IbanTemplate) {
	f, err := country.NewIbanTemplate(template)
	if !assert.NoError(t, err, "failed to parse iban template, %s", err) {
		return
	}

	return *f
}

func TestCountriesFormat(t *testing.T) {
	db, err := NewCountriesFormat("../../test-data/countries-test.yaml")
	if !assert.NoError(t, err, "failed to init db") {
		return
	}
	defer db.Shutdown()

	expected := map[string]*country.Format{
		"BA": {
			Country:  "Bosnia and Herzegovina",
			Len:      20,
			Template: testParseTemplate(t, "16n"),
			Kk:       "39",
		},
		"BR": {
			Country:  "Brazil",
			Template: testParseTemplate(t, "23n,1a,1c"),
			Len:      29,
		},
		"CY": {
			Country:  "Cyprus",
			Template: testParseTemplate(t, "8n,16c"),
			Len:      28,
		},
	}

	assert.Equal(t, expected, db.records)

	countryCode := "BA"
	f, err := db.Get("BA")
	if assert.NoError(t, err, "failed to get a format record by country code ", countryCode) {
		assert.Equal(t, expected[countryCode], f)
	}

	_, err = db.Get("UNK")
	assert.Error(t, err, "failed to catch an error for an unknown country code")
}

func TestCountriesFormat_Invalid(t *testing.T) {
	db, err := NewCountriesFormat("../../unknown-data/unknown.unk")
	assert.Error(t, err, "failed to catch an error")
	db.Shutdown()
}
