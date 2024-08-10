package logger

import (
	"io"

	"gopkg.in/natefinch/lumberjack.v2"
)

type FileWriter interface {
	io.Writer
}

func NewFileWriter(filename string, maxSize int, maxBackups int, compress bool) FileWriter {
	fileWriter := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		Compress:   compress,
	}

	return fileWriter
}
