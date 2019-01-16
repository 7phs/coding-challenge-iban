package helper

import (
	"os"
	"strconv"
	"strings"
)

func EnvStr(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}

	return val
}

func EnvInt(key string, def int) int {
	str := os.Getenv(key)
	if str == "" {
		return def
	}

	val, err := strconv.Atoi(str)
	if err != nil {
		return def
	}

	return val
}

func EnvBool(key string, def bool) bool {
	str := os.Getenv(key)
	if str == "" {
		return def
	}

	switch strings.ToLower(str) {
	case "true":
		return true
	case "false":
		return false
	default:
		return def
	}
}
