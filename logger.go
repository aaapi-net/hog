package hog

import (
	"log"
	"os"
	"sync"
)

type LogLevel int

const (
	LogLevelNone LogLevel = iota
	LogLevelError
	LogLevelInfo
	LogLevelDebug
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	IsEnabled(level LogLevel) bool
}

type defaultLogger struct {
	level  LogLevel
	logger *log.Logger
	mu     sync.Mutex
}

func newDefaultLogger(level LogLevel) *defaultLogger {
	return &defaultLogger{
		level:  level,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (l *defaultLogger) Debug(args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level >= LogLevelDebug {
		l.logger.Print(append([]interface{}{"[DEBUG]"}, args...)...)
	}
}

func (l *defaultLogger) Info(args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level >= LogLevelInfo {
		l.logger.Print(append([]interface{}{"[INFO]"}, args...)...)
	}
}

func (l *defaultLogger) Error(args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level >= LogLevelError {
		l.logger.Print(append([]interface{}{"[ERROR]"}, args...)...)
	}
}

func (l *defaultLogger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

func (l *defaultLogger) GetLevel() LogLevel {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}

func (l *defaultLogger) IsEnabled(level LogLevel) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level >= level
}
