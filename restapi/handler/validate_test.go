package handler

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/verdverm/frisby"
)

func TestValidate(t *testing.T) {
	defer testGinMode()()

	srv, df := testDefaultRouter(t)
	defer df()

	testSuites := []*struct {
		Description string
		Iban        string
		Status      int
		Content     string
	}{
		{
			Description: "valid",
			Iban:        "BR 97 00360305 00001 0009795493 P 1",
			Status:      http.StatusOK,
			Content:     `{"iban":"valid"}`,
		},
		{
			Description: "invalid",
			Iban:        "BR 97 289312873612786",
			Status:      http.StatusPreconditionFailed,
			Content:     `{"iban":"invalid"}`,
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
