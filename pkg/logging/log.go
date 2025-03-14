package logging

import (
	"github.com/sirupsen/logrus"
)

type Level = logrus.Level

const (
	PanicLevel Level = logrus.PanicLevel
	FatalLevel Level = logrus.FatalLevel
	ErrorLevel Level = logrus.ErrorLevel
	WarnLevel  Level = logrus.WarnLevel
	InfoLevel  Level = logrus.InfoLevel
	DebugLevel Level = logrus.DebugLevel
	TraceLevel Level = logrus.TraceLevel
)

type Logger struct {
	*logrus.Entry
}

func NewLogger(level Level) (*Logger, error) {
	l := logrus.New()
	l.Level = level
	return &Logger{
		Entry: logrus.NewEntry(l),
	}, nil
}

func ParseLevel(level string) (Level, error) {
	return logrus.ParseLevel(level)
}

func (l *Logger) WithField(field string, value any) *Logger {
	return &Logger{
		Entry: l.Entry.WithField(field, value),
	}
}
