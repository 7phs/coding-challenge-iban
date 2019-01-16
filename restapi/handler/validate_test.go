package handler

import (
	"github.com/7phs/coding-challenge-iban/config"
	"github.com/stretchr/testify/assert"
	"github.com/verdverm/frisby"
	"net/http/httptest"
	"testing"
)

func TestValidate(t *testing.T) {
	defer testGinMode()()

	srv := httptest.NewServer(DefaultRouter(&config.Config{}))
	defer srv.Close()

	testSuites := []*struct {
		Description string
		Iban        string
		Status      int
		Content     string
	}{
		{

		},
	}

	for _, test := range testSuites {
		f := frisby.Create("Test validate: " + test.Description).
			Get(srv.URL + "/validate/" + test.Iban).
			Send().
			ExpectStatus(test.Status).
			ExpectContent(test.Content)

		assert.Empty(t, f.Errs, "%s", f.Errs)
	}
}
