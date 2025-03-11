package hog

import (
	"log"
	"os"
	"sync"
)

// LogLevel defines the verbosity of logging.
type LogLevel int

const (
	LogLevelNone LogLevel = iota
	LogLevelError
	LogLevelInfo
	LogLevelDebug
)

// Logger defines methods for request/response logging.
type Logger interface {
	// Debug logs debug level information.
	Debug(args ...interface{})

	// Info logs informational messages.
	Info(args ...interface{})

	// Error logs error information.
	Error(args ...interface{})

	// IsEnabled checks if a specific log level is active.
	IsEnabled(level LogLevel) bool
}

// defaultLogger implements Logger with standard log package.
type defaultLogger struct {
	level  LogLevel
	logger *log.Logger
	mu     sync.Mutex
}

// newDefaultLogger creates a new default logger with specified logging level.
func newDefaultLogger(level LogLevel) *defaultLogger {
	return &defaultLogger{
		level:  level,
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

// Debug logs debug level information.
func (l *defaultLogger) Debug(args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level >= LogLevelDebug {
		l.logger.Print(append([]interface{}{"[DEBUG]"}, args...)...)
	}
}

// Info logs informational messages.
func (l *defaultLogger) Info(args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level >= LogLevelInfo {
		l.logger.Print(append([]interface{}{"[INFO]"}, args...)...)
	}
}

// Error logs error information.
func (l *defaultLogger) Error(args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.level >= LogLevelError {
		l.logger.Print(append([]interface{}{"[ERROR]"}, args...)...)
	}
}

// SetLevel changes the current logging level.
func (l *defaultLogger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// GetLevel returns the current logging level.
func (l *defaultLogger) GetLevel() LogLevel {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level
}

// IsEnabled checks if a specific log level is active.
func (l *defaultLogger) IsEnabled(level LogLevel) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.level >= level
}
