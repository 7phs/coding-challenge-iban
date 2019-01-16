package handler

import (
	"github.com/7phs/coding-challenge-iban/config"
	"github.com/7phs/coding-challenge-iban/restapi/common"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func DefaultRouter(conf *config.Config) http.Handler {
	log.Info("http/router: init")

	router := gin.New()

	router.Use(gin.Logger())
	if conf.Http().Cors() {
		router.Use(AllowCors())
	}
	// VALIDATE
	router.GET("/validate/:IBAN", common.NewHandler(conf, &ValidateHandler{}))
	// HEALTH CHECK
	router.GET("/health/check", HealthCheck)

	return router
}
