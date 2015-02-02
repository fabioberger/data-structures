package utils

import (
	"io"
	"log"
)

func SetLoggerOut(logger *log.Logger, out io.Writer) *log.Logger {
	return log.New(out, "", 0)
}
