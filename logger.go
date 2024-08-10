package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	writer io.Writer
}

type LoggerConfig struct {
	ConsoleWriter bool
	FileWriter
}

func New(loggerConfig *LoggerConfig) *Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var writers []io.Writer

	if loggerConfig.ConsoleWriter {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if loggerConfig.FileWriter != nil {
		writers = append(writers, loggerConfig.FileWriter)
	}

	multiWriter := io.MultiWriter(writers...)

	return &Logger{multiWriter}
}

func (l *Logger) Sub(topic string) *SubLogger {
	logger := zerolog.New(l.writer)
	return &SubLogger{topic, &logger}
}
