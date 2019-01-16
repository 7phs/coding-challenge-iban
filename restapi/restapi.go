package restapi

import (
	"github.com/7phs/coding-challenge-iban/config"
	"github.com/gin-gonic/gin"
)

func Init(conf *config.Config) {
	switch conf.Stage() {
	case config.StageTest:
		gin.SetMode(gin.TestMode)
	case config.StageDev:
		gin.SetMode(gin.DebugMode)
	case config.StageProd:
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
