package handler

import (
	"github.com/7phs/coding-challenge-iban/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllowCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	}
}

func Options(conf *config.Config, methods string) func(*gin.Context) {
	if conf.Http().Cors() {
		return func(c *gin.Context) {
			c.Header("Access-Control-Allow-Methods", methods)
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Headers", "accept, content-type")
			c.String(http.StatusOK, "ok")
		}
	} else {
		return func(c *gin.Context) {
			c.String(http.StatusOK, "ok")
		}
	}
}
