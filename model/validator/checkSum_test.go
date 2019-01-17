package validator

import (
	"testing"

	"github.com/7phs/coding-challenge-iban/model/records"
	"github.com/stretchr/testify/assert"
)

func TestCheckSum(t *testing.T) {
	testSuites := []*struct {
		in  string
		exp bool
	}{
		{
			in:  "GB82 WEST 1234 5698 7654 32",
			exp: true,
		},
		{
			in:  "BR 97 00360305 00001 0009795493 P 1",
			exp: true,
		},
		{
			in:  "AA123902347891287612378",
			exp: false,
		},
	}

	validator := NewCheckSum()

	for i, test := range testSuites {
		err := validator.Validate(records.NewIban(test.in))
		if !test.exp {
			assert.Error(t, err, "%d: failed to catch error", i)
		} else {
			assert.NoError(t, err, "%d: wrong validate and got %v", i, err)
		}
	}
}
