package country

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIbanTemplate(t *testing.T) {
	testSuites := []*struct {
		in  string
		exp string
		err bool
	}{
		{
			in:  "",
			err: true,
		},
		{
			in:  "7n",
			exp: "^[0-9]{7}$",
		},
		{
			in:  "19a",
			exp: "^[A-Z]{19}$",
		},
		{
			in:  "2c",
			exp: "^[a-zA-Z0-9]{2}$",
		},
		{
			in:  "23n,1a,1c",
			exp: "^[0-9]{23}[A-Z]{1}[a-zA-Z0-9]{1}$",
		},
		{
			in:  "23z,1a,1c",
			err: true,
		},
		{
			in:  "14a,15n,1za",
			err: true,
		},
	}

	for _, test := range testSuites {
		exist, err := ParseIbanTemplate(test.in)
		if test.err {
			assert.Error(t, err, "not catch a parsing error")
		} else if assert.NoError(t, err, "got a parsing error, %v", err) {
			assert.Equal(t, test.exp, exist)
		}
	}
}
