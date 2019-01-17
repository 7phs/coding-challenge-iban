package handler

import (
	"net/http"

	"github.com/7phs/coding-challenge-iban/config"
	"github.com/7phs/coding-challenge-iban/model"
	"github.com/7phs/coding-challenge-iban/restapi/common"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func DefaultRouter(conf *config.Config, models model.Models) http.Handler {
	log.Info("http/router: init")

	router := gin.New()

	router.Use(gin.Logger())
	// VALIDATE
	router.GET("/validate/:IBAN", common.NewHandler(conf, models, &ValidateHandler{}))
	// HEALTH CHECK
	router.GET("/health/check", HealthCheck)

	return router
}
