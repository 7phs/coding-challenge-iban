package config

import (
	"fmt"
	"os"

	"github.com/7phs/coding-challenge-iban/helper"
	log "github.com/sirupsen/logrus"
)

const (
	defStage      = StageDev
	defLogLevel   = LogLevelDebug
	defAddress    = ":8080"
	defDbPath     = "./data/countries-iban.yaml"
	defTextLength = 1024
	defCors       = false

	EnvStage      = "STAGE"
	EnvLogLevel   = "LOG_LEVEL"
	EnvAddress    = "ADDRESS"
	EnvDbPath     = "DB_PATH"
	EnvCors       = "CORS"
	EnvTextLength = "LIMIT_TEXT_LENGTH"
)

type Config struct {
	stage    Stage
	logLevel LogLevel

	address string
	dbPath  string

	http  Http
	limit Limit
}

func ParseConfig() *Config {
	return &Config{
		stage:    NewStage(helper.EnvStr(EnvStage, defStage.String())),
		logLevel: NewLogLevel(helper.EnvStr(EnvLogLevel, defLogLevel.String())),

		address: helper.EnvStr(EnvAddress, defAddress),
		dbPath:  helper.EnvStr(EnvDbPath, defDbPath),

		http: Http{
			cors: helper.EnvBool(EnvCors, defCors),
		},

		limit: Limit{
			textLength: helper.EnvInt(EnvTextLength, defTextLength),
		},
	}
}

func (o *Config) LogLevel() LogLevel {
	return o.logLevel
}

func (o *Config) Stage() Stage {
	return o.stage
}

func (o *Config) Address() string {
	return o.address
}

func (o *Config) DbPath() string {
	return o.dbPath
}

func (o *Config) Http() *Http {
	return &o.http
}

func (o *Config) Limit() *Limit {
	return &o.limit
}

func (o *Config) Validate() error {
	var errList helper.ErrList

	if len(o.dbPath) == 0 {
		errList.Add(fmt.Errorf(EnvDbPath + ": empty"))
	} else if info, err := os.Stat(o.dbPath); err != nil || info.IsDir() {
		if err != nil {
			errList.Add(fmt.Errorf(EnvDbPath+": %v", err))
		} else {
			errList.Add(fmt.Errorf(EnvDbPath + ": is a directory, but wait a file"))
		}
	}

	if len(o.address) == 0 {
		errList.Add(fmt.Errorf(EnvAddress + ": empty"))
	}

	if o.stage == StageUnknown {
		errList.Add(fmt.Errorf(EnvStage+": unsupported, support - %s", StageAll))
	}

	if o.logLevel == LogLevelUnknown {
		errList.Add(fmt.Errorf(EnvLogLevel+": unsupported, support - %s", LogLevelAll))
	}

	return errList.Result()
}

func (o *Config) Dump() {
	log.Info("config: stage - ", o.stage)
	log.Info("config: countries db path - ", o.dbPath)
}

type Limit struct {
	textLength int
}

func (o *Limit) TextLength() int {
	return o.textLength
}

type Http struct {
	cors bool
}

func (o *Http) Cors() bool {
	return o.cors
}
