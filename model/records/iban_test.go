package records

import (
	"strings"
	"testing"

	"github.com/bmizerany/assert"
)

func TestIban(t *testing.T) {
	var (
		iban        = "BR 97 00360305 00001 0009795493 P 1"
		text        = strings.Replace(iban, " ", "", -1)
		countryCode = "BR"
		kk          = "97"
		suffix      = text[4:]
		number      = "360305000010009795493251112797"
		formatted   = "BR97 0036 0305 0000 1000 9795 493P 1"
	)

	rec := NewIban(iban)

	assert.Equal(t, rec.Raw(), iban, "raw")
	assert.Equal(t, rec.Text(), text, "text")
	assert.Equal(t, rec.Len(), len(text), "length")
	assert.Equal(t, rec.CountryCode(), countryCode, "country code")
	assert.Equal(t, rec.Kk(), kk, "kk")
	assert.Equal(t, rec.Suffix(), suffix, "suffix")
	assert.Equal(t, rec.String(), formatted, "string")
	assert.Equal(t, rec.Number().String(), number, "number")
}

func TestIbanInvalid(t *testing.T) {
	var (
		iban        = "<>)( WE..S!!!! |||||"
		text        = "----"
		countryCode = "--"
		kk          = "--"
		suffix      = ""
		number      = "0"
	)

	rec := NewIban(iban)

	assert.Equal(t, rec.Raw(), iban, "raw")
	assert.Equal(t, rec.Text(), text, "text")
	assert.Equal(t, rec.Len(), len(text), "length")
	assert.Equal(t, rec.CountryCode(), countryCode, "country code")
	assert.Equal(t, rec.Kk(), kk, "kk")
	assert.Equal(t, rec.Suffix(), suffix, "suffix")
	assert.Equal(t, rec.String(), text, "string")
	assert.Equal(t, rec.Number().String(), number, "number")
}
