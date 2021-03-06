package country

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
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

	for i, test := range testSuites {
		exist, err := ParseIbanTemplate(test.in)
		if test.err {
			assert.Error(t, err, "%d: not catch a parsing error", i)
		} else if assert.NoError(t, err, "%d: got a parsing error, %v", i, err) {
			assert.Equal(t, test.exp, exist)
		}

		if len(test.in) > 0 {
			obj := &IbanTemplate{}
			err = yaml.Unmarshal([]byte(test.in), &obj)
			if test.err {
				assert.Error(t, err, "%d: not catch a parsing error", i)
			} else if assert.NoError(t, err, "%d: got a parsing error, %v", i, err) {
				assert.Equal(t,
					&IbanTemplate{
						Regexp: regexp.MustCompile(test.exp),
					},
					obj)
			}
		}
	}
}
