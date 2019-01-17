package country

import (
	"testing"

	"github.com/7phs/coding-challenge-iban/helper"
	"github.com/stretchr/testify/assert"
)

func TestFormat(t *testing.T) {
	testSuites := []*struct {
		kk, text, suffix string
		in               *Format
		exp              error
	}{
		{
			kk:     "98",
			text:   "GE987890HJKA",
			suffix: "7890HJKA",
			in: &Format{
				Len:      12,
				Template: *MustIbanTemplate("4n,4a"),
				Kk:       "98",
			},
		},
		{
			kk:     "11",
			text:   "GE117890Hj8AHUIU",
			suffix: "7890Hj8AHUIU",
			in: &Format{
				Len:      16,
				Template: *MustIbanTemplate("4n,4c,4a"),
			},
		},
		{
			kk:     "11",
			text:   "GE117890Hj8AHUIU",
			suffix: "7890Hj8AHUIU",
			in: &Format{
				Len:      20,
				Template: *MustIbanTemplate("4n,4a,4a"),
			},
			exp: helper.ErrList{
				ErrCountryLength,
				ErrCountryTemplate,
			},
		},
		{
			kk:     "11",
			text:   "GE117890Hj8AHUIU",
			suffix: "7890Hj8AHUIU",
			in: &Format{
				Len:      16,
				Template: *MustIbanTemplate("4n,4c,4a"),
				Kk:       "07",
			},
			exp: helper.ErrList{
				ErrCountryKk,
			},
		},
		{
			kk:     "11",
			text:   "GE117890Hj8AHUIU",
			suffix: "7890Hj8AHUIU",
			in: &Format{
				Len: 16,
				Kk:  "11",
			},
			exp: helper.ErrList{
				ErrCountryEmptyTemplate,
			},
		},
	}

	for i, test := range testSuites {
		assert.Equal(t, test.exp, test.in.Validate(test.kk, test.text, test.suffix), "%d", i)
	}
}
