package config

import (
	"fmt"
	"strings"
)

const (
	StageProd    Stage = "prod"
	StageDev     Stage = "dev"
	StageTest    Stage = "test"
	StageUnknown Stage = "unknown"
)

var (
	StageAll = strings.Join([]string{
		string(StageProd),
		string(StageDev),
		string(StageTest),
	}, ", ")
)

type Stage string

func NewStage(str string) Stage {
	switch strings.ToLower(str) {
	case "prod":
		return StageProd
	case "dev":
		return StageDev
	case "test":
		return StageTest
	default:
		return StageUnknown
	}
}

func (o Stage) String() string {
	switch o {
	case StageProd:
		return "prod"
	case StageDev:
		return "dev"
	case StageTest:
		return "test"
	case StageUnknown:
		return "unknown"
	default:
		return fmt.Sprintf("[INVALID: %s]", string(o))
	}
}
