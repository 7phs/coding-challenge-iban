package validator

import (
	"os"
	"testing"

	"github.com/7phs/coding-challenge-iban/model/records"
	"github.com/stretchr/testify/assert"
)

type testValidator struct {
	valid bool
}

func (o *testValidator) Validate(*records.Iban) error {
	if o.valid {
		return nil
	}

	return os.ErrInvalid
}

func TestFlow(t *testing.T) {
	err := NewFlow().
		Then(&testValidator{valid: true}).
		Then(&testValidator{valid: true}).
		Validate(&records.Iban{})
	assert.NoError(t, err, "valid")

	err = NewFlow().
		Then(&testValidator{valid: true}).
		Then(&testValidator{valid: false}).
		Then(&testValidator{valid: true}).
		Validate(&records.Iban{})
	assert.Error(t, err, "invalid")

	err = NewFlow().
		Then(&testValidator{valid: true}).
		Then(&testValidator{valid: true}).
		Then(&testValidator{valid: false}).
		Validate(&records.Iban{})
	assert.Error(t, err, "invalid")
}
