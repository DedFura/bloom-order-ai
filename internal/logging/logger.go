package logging

import (
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Logger interface {
	SetLevel(level string) error
	Debug(args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	WithFields(fields map[string]interface{}) *TemporaryFields
	WithError(err error) *TemporaryFields
}

type LogrusLogger struct {
	*logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	return &LogrusLogger{
		Logger: logger,
	}
}

func (l *LogrusLogger) SetLevel(level string) error {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return errors.Wrap(err, "failed to parse logger")
	}
	l.Logger.SetLevel(lvl)

	return nil
}

func (l *LogrusLogger) Debug(args ...interface{}) {
	l.Logger.Debug(args...)
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.Logger.Info(args...)
}

func (l *LogrusLogger) Infof(format string, args ...interface{}) {
	l.Logger.Infof(format, args...)
}

func (l *LogrusLogger) Warn(args ...interface{}) {
	l.Logger.Warn(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.Logger.Error(args...)
}

func (l *LogrusLogger) Fatal(args ...interface{}) {
	l.Logger.Fatal(args...)
}

func (l *LogrusLogger) WithFields(fields map[string]interface{}) *TemporaryFields {
	return &TemporaryFields{
		Logger: l,
		Entry:  l.Logger.WithFields(fields),
	}
}

func (l *LogrusLogger) WithError(err error) *TemporaryFields {
	return &TemporaryFields{
		Logger: l,
		Entry:  l.Logger.WithError(err),
	}
}

type TemporaryFields struct {
	Logger *LogrusLogger
	Entry  *logrus.Entry
}

func (t *TemporaryFields) Info(args ...interface{}) {
	t.Entry.Info(args...)
}

func (t *TemporaryFields) Warn(args ...interface{}) {
	t.Entry.Warn(args...)
}

func (t *TemporaryFields) Error(args ...interface{}) {
	t.Entry.Error(args...)
}

func (t *TemporaryFields) WithError(err error) *TemporaryFields {
	return &TemporaryFields{
		Logger: t.Logger,
		Entry:  t.Entry.WithError(err),
	}
}
