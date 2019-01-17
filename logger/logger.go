package logger

import (
	"bytes"
	"log"
	"os"

	"github.com/7phs/coding-challenge-iban/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	defLogLevel = logrus.DebugLevel
)

var (
	logLevelMap = map[config.LogLevel]logrus.Level{
		config.LogLevelDebug:   logrus.DebugLevel,
		config.LogLevelInfo:    logrus.InfoLevel,
		config.LogLevelWarning: logrus.WarnLevel,
		config.LogLevelError:   logrus.ErrorLevel,
	}
)

type logWriter struct{}

func (o *logWriter) Write(p []byte) (n int, err error) {
	p = bytes.Replace(p, []byte{'\n'}, []byte{' '}, -1)

	logrus.Print(string(p))
	return len(p), nil
}

func Init(conf *config.Config) {
	// gin specific
	log.SetOutput(&logWriter{})
	gin.DefaultWriter = &logWriter{}
	gin.DefaultErrorWriter = &logWriter{}
	// init logger
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	// set a log level
	logLevel, ok := logLevelMap[conf.LogLevel()]
	if !ok {
		logLevel = defLogLevel
	}
	logrus.SetLevel(logLevel)
}
