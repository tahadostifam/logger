package main

import (
	"errors"

	"github.com/tahadostifam/logger"
)

func main() {
	logger := logger.New(&logger.LoggerConfig{
		ConsoleWriter: true,
		FileWriter:    logger.NewFileWriter("./logs/logs", 2, 1, true),
	})

	fooLogger := logger.Sub("Foo")
	barLogger := logger.Sub("Bar")

	fooLogger.Error("unable to create the foo in database", errors.New("error detail"), map[string]string{
		"name": "Sample key for field of error method",
	})

	fooLogger.Warn("unable to create the foo in database", map[string]string{
		"name": "Sample key for field of warn method",
	})

	fooLogger.Info("hello world hello world hello world", map[string]string{
		"name": "Sample key for field of warn method",
	})

	barLogger.Debug("debug mebug debug mebug debug mebug", map[string]string{
		"name": "Sample key for field of warn method",
	})

	barLogger.Trace("trace mrace trace mrace trace mrace", map[string]string{
		"name": "Sample key for field of warn method",
	})
}
