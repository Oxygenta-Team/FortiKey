package logging

import "github.com/sirupsen/logrus"

type Logger struct {
	*logrus.Entry
}

func NewLogger(level logrus.Level) *Logger {
	l := logrus.New()
	l.Level = level

	return &Logger{
		Entry: logrus.NewEntry(l),
	}
}

func (l *Logger) WithField(field string, value any) *Logger {
	return &Logger{
		Entry: l.Entry.WithField(field, value),
	}
}
