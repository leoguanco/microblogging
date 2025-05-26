package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	entry *logrus.Entry
}

type Fields = logrus.Fields

var log = logrus.New()

func init() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.JSONFormatter{})

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.SetLevel(logrus.InfoLevel)
	} else {
		log.SetLevel(level)
	}
}

func GetLogger() *Logger {
	return &Logger{
		entry: logrus.NewEntry(log),
	}
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{
		entry: l.entry.WithField(key, value),
	}
}

func (l *Logger) WithFields(fields Fields) *Logger {
	return &Logger{
		entry: l.entry.WithFields(logrus.Fields(fields)),
	}
}

func (l *Logger) WithError(err error) *Logger {
	return &Logger{
		entry: l.entry.WithError(err),
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}
