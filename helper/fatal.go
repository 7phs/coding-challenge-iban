package helper

import (
	"fmt"
	"os"
)

func Fatal(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(-1)
}
