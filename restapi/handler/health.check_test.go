package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/verdverm/frisby"
)

func TestHealthCheck(t *testing.T) {
	defer testGinMode()()

	srv, df := testDefaultRouter(t)
	defer df()

	f := frisby.Create("Test successful a health check").
		Get(srv.URL + "/health/check").
		Send().
		ExpectStatus(200).
		ExpectContent("ok")

	assert.Empty(t, f.Errs, "%s", f.Errs)
}
