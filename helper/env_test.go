package helper

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupEnv(key string) func() {
	v := os.Getenv(key)

	return func() {
		os.Setenv(key, v)
	}
}

func TestEnvStr(t *testing.T) {
	key := "--TST--"

	defer setupEnv(key)()

	testSuites := []*struct {
		in  string
		def string
		exp string
	}{
		{
			in:  "1234",
			def: "",
			exp: "1234",
		},
		{
			in:  "abc",
			def: "000",
			exp: "abc",
		},
		{
			in:  "",
			def: "000",
			exp: "000",
		},
	}

	for _, test := range testSuites {
		os.Setenv(key, test.in)

		assert.Equal(t, test.exp, EnvStr(key, test.def))
	}
}

func TestEnvInt(t *testing.T) {
	key := "--TST--"

	defer setupEnv(key)()

	testSuites := []*struct {
		in  string
		def int
		exp int
	}{
		{
			in:  "1234",
			exp: 1234,
		},
		{
			in:  "23abc",
			exp: 0,
		},
		{
			def: 9193,
			exp: 9193,
		},
	}

	for _, test := range testSuites {
		os.Setenv(key, test.in)

		assert.Equal(t, test.exp, EnvInt(key, test.def))
	}
}

func TestEnvBool(t *testing.T) {
	key := "--TST--"

	defer setupEnv(key)()

	testSuites := []*struct {
		in  string
		def bool
		exp bool
	}{
		{
			in:  "true",
			exp: true,
		},
		{
			in:  "false",
			exp: false,
		},
		{
			in:  "unknown",
			exp: false,
		},
		{
			def: true,
			exp: true,
		},
	}

	for _, test := range testSuites {
		os.Setenv(key, test.in)

		assert.Equal(t, test.exp, EnvBool(key, test.def))
	}
}
