package validator

import (
	"testing"

	"github.com/7phs/coding-challenge-iban/model/records"
	"github.com/stretchr/testify/assert"
)

func TestRaw(t *testing.T) {
	testSuites := []*struct {
		in  string
		exp bool
	}{
		{
			in:  "GB82WEST12345698765432",
			exp: true,
		},
		{
			in:  "1239023.? - $ 47891287612378 GHJG",
			exp: false,
		},
	}

	validator := NewRaw()

	for i, test := range testSuites {
		err := validator.Validate(records.NewIban(test.in))
		if !test.exp {
			assert.Error(t, err, "%d: failed to catch error", i)
		} else {
			assert.NoError(t, err, "%d: wrong validate and got %v", i, err)
		}
	}

}
