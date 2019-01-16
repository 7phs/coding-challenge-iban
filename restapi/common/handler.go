package common

import (
	"github.com/7phs/coding-challenge-iban/config"
	"github.com/7phs/coding-challenge-iban/helper"
	"github.com/7phs/coding-challenge-iban/model/errCode"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type HandlerFactory interface {
	New() HandlerImpl
}

type HandlerImpl interface {
	LogPrefix() string
	Bind(c *gin.Context) error
	Validate(conf *config.Config) error
	Process(c *gin.Context)
}

func NewHandler(conf *config.Config, factory HandlerFactory) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler := factory.New()

		// BIND PARAMS
		if err := handler.Bind(c); err != nil {
			log.Error(handler.LogPrefix(), ": failed to bind parameters - ", err)

			c.JSON(http.StatusBadRequest, helper.NewGeneralErrorResponse(errCode.ErrParamBinding, err))
			return
		}
		// VALIDATE PARAMS (empty, invalid format, etc.)
		if err := handler.Validate(conf); err != nil {
			log.Error(handler.LogPrefix(), ": failed to validate parameters - ", err)

			c.JSON(http.StatusUnprocessableEntity, helper.NewGeneralErrorResponse(errCode.ErrParamValidation, err))
			return
		}

		handler.Process(c)
	}
}
