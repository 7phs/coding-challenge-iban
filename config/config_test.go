package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupCofigTest() func() {
	prev := map[string]string{}

	for _, name := range []string{
		EnvStage, EnvLogLevel,
		EnvAddress, EnvDbPath,
		EnvTextLength,
	} {
		prev[name] = os.Getenv(name)
	}

	return func() {
		for name, value := range prev {
			os.Setenv(name, value)
		}
	}
}

func testSetEnv(env map[string]string) {
	for key, value := range env {
		os.Setenv(key, value)
	}
}

func resetEnv() {
	for _, name := range []string{
		EnvStage, EnvLogLevel,
		EnvAddress, EnvDbPath,
		EnvTextLength,
	} {
		os.Setenv(name, "")
	}
}

func defConfig(config *Config) *Config {
	if config.stage == "" {
		config.stage = defStage
	}
	if config.logLevel == "" {
		config.logLevel = defLogLevel
	}
	if config.address == "" {
		config.address = defAddress
	}
	if config.dbPath == "" {
		config.dbPath = defDbPath
	}
	if config.limit.textLength == 0 {
		config.limit.textLength = defTextLength
	}

	return config
}

func TestConfig(t *testing.T) {
	defer setupCofigTest()()

	existingDbPath := "../test-data/countries-test.yaml"

	testSuites := []*struct {
		env map[string]string
		exp *Config
		err bool
	}{
		{
			env: map[string]string{
				EnvStage:      StageDev.String(),
				EnvLogLevel:   LogLevelWarning.String(),
				EnvAddress:    ":8080",
				EnvDbPath:     existingDbPath,
				EnvTextLength: "4012",
			},
			exp: &Config{
				stage:    StageDev,
				logLevel: LogLevelWarning,

				address: ":8080",
				dbPath:  "../test-data/countries-test.yaml",

				limit: Limit{
					textLength: 4012,
				},
			},
		},
		{
			env: map[string]string{
				EnvStage:    StageTest.String(),
				EnvLogLevel: LogLevelInfo.String(),
				EnvDbPath:   existingDbPath,
			},
			exp: defConfig(&Config{
				stage:    StageTest,
				logLevel: LogLevelInfo,
				dbPath:   existingDbPath,
			}),
		},
		{
			env: map[string]string{
				EnvStage:    StageProd.String(),
				EnvLogLevel: LogLevelDebug.String(),
				EnvDbPath:   existingDbPath,
			},
			exp: defConfig(&Config{
				stage:    StageProd,
				logLevel: LogLevelDebug,
				dbPath:   existingDbPath,
			}),
		},
		{
			env: map[string]string{
				EnvLogLevel: LogLevelError.String(),
				EnvDbPath:   existingDbPath,
			},
			exp: defConfig(&Config{
				logLevel: LogLevelError,
				dbPath:   existingDbPath,
			}),
		},
		{
			env: map[string]string{
				EnvStage: StageUnknown.String(),
			},
			exp: defConfig(&Config{
				stage: Stage("unknown"),
			}),
			err: true,
		},
		{
			env: map[string]string{
				EnvLogLevel: LogLevelUnknown.String(),
			},
			exp: defConfig(&Config{
				logLevel: LogLevelUnknown,
			}),
			err: true,
		},
		{
			env: map[string]string{
				EnvDbPath: "unknown/../unknown",
			},
			exp: defConfig(&Config{
				dbPath: "unknown/../unknown",
			}),
			err: true,
		},
		{
			env: map[string]string{
				EnvDbPath: ".",
			},
			exp: defConfig(&Config{
				dbPath: ".",
			}),
			err: true,
		},
	}

	for _, test := range testSuites {
		resetEnv()
		testSetEnv(test.env)

		exist := ParseConfig()
		assert.Equal(t, test.exp, exist)
		assert.Equal(t, test.exp.stage, exist.Stage())
		assert.Equal(t, test.exp.logLevel, exist.LogLevel())
		assert.Equal(t, test.exp.address, exist.Address())
		assert.Equal(t, test.exp.dbPath, exist.DbPath())
		assert.Equal(t, test.exp.limit.textLength, exist.Limit().TextLength())

		err := exist.Validate()
		if test.err {
			assert.Error(t, err, "failed to catch an error")
		} else {
			assert.NoError(t, err, "invalid config")
		}
	}
}
