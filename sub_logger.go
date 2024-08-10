package logger

import (
	"reflect"

	"github.com/rs/zerolog"
)

type SubLogger struct {
	topic  string
	logger *zerolog.Logger
}

type Fields = map[string]string

func addFields(fields Fields, event *zerolog.Event) {
	val := reflect.ValueOf(fields)
	keys := val.MapKeys()

	for _, v := range keys {
		event.Str(v.String(), fields[v.String()])
	}
}

func (sl *SubLogger) Panic(reason string, err error, fields Fields) {
	e := sl.logger.Panic()
	e.Timestamp()
	e.Err(err)
	e.Str("topic", sl.topic)
	addFields(fields, e)
	e.Msg(reason)
}

func (sl *SubLogger) Error(reason string, err error, fields Fields) {
	e := sl.logger.Error()
	e.Timestamp()
	e.Err(err)
	e.Str("topic", sl.topic)
	addFields(fields, e)
	e.Msg(reason)
}

func (sl *SubLogger) Fatal(reason string, err error, fields Fields) {
	e := sl.logger.Fatal()
	e.Timestamp()
	e.Err(err)
	e.Str("topic", sl.topic)
	addFields(fields, e)
	e.Msg(reason)
}

func (sl *SubLogger) Warn(msg string, fields Fields) {
	e := sl.logger.Warn()
	e.Timestamp()
	e.Str("topic", sl.topic)
	addFields(fields, e)
	e.Msg(msg)
}

func (sl *SubLogger) Info(msg string, fields Fields) {
	e := sl.logger.Info()
	e.Timestamp()
	e.Str("topic", sl.topic)
	addFields(fields, e)
	e.Msg(msg)
}

func (sl *SubLogger) Debug(detail string, fields Fields) {
	e := sl.logger.Debug()
	e.Timestamp()
	e.Str("topic", sl.topic)
	addFields(fields, e)
	e.Msg(detail)
}

func (sl *SubLogger) Trace(detail string, fields Fields) {
	e := sl.logger.Trace()
	e.Timestamp()
	e.Str("topic", sl.topic)
	addFields(fields, e)
	e.Msg(detail)
}
