package config

import (
	"github.com/7phs/coding-challenge-iban/helper"
)

const (
	defStage      = StageDev
	defLogLevel   = LogLevelDebug
	defAddress    = ":8080"
	defDbPath     = "./data/countries-iban.yaml"
	defTextLength = 1024
	defCors       = false
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
		stage:    NewStage(helper.EnvStr("STAGE", defStage.String())),
		logLevel: NewLogLevel(helper.EnvStr("LOG_LEVEL", defLogLevel.String())),

		address: helper.EnvStr("ADDRESS", defAddress),
		dbPath:  helper.EnvStr("DB_PATH", defDbPath),

		http: Http{
			cors: helper.EnvBool("CORS", defCors),
		},

		limit: Limit{
			textLength: helper.EnvInt("LIMIT_TEXT_LENGTH", defTextLength),
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
	return nil
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
