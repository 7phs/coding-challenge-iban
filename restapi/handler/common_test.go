package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/7phs/coding-challenge-iban/config"
	"github.com/7phs/coding-challenge-iban/db/file"
	"github.com/7phs/coding-challenge-iban/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func testGinMode() func() {
	mode := gin.Mode()
	gin.SetMode(gin.TestMode)

	return func() {
		gin.SetMode(mode)
	}
}

func testDefaultRouter(t *testing.T) (*httptest.Server, func()) {
	db, err := file.NewCountriesFormat("../../test-data/countries-test.yaml")
	assert.NoError(t, err, "failed to init db")

	srv := httptest.NewServer(DefaultRouter(&config.Config{},
		&model.DefaultState{
			Valid: model.DefaultValidator(&model.Dependencies{
				CountriesFormat: db,
			}),
		}))

	return srv, func() {
		srv.Close()
		db.Shutdown()
	}
}
