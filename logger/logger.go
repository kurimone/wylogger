package logger

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"
	"sync"
	"time"
)

// LogLevel defines the severity of the log message.
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// ANSI color codes
const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
	colorReset  = "\033[0m"
)

// Map of log level to color
var logLevelColors = map[LogLevel]string{
	DEBUG: colorBlue,
	INFO:  colorGreen,
	WARN:  colorYellow,
	ERROR: colorRed,
	FATAL: colorPurple,
}

// Logger wraps the standard log.Logger from the Go standard library.
type Logger struct {
	*log.Logger
	level LogLevel
}

// Singleton instance of Logger
var (
	globalLogger *Logger
	once         sync.Once
)

// NewLogger creates a new Logger instance with the specified log level.
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", 0), // Disable the default timestamp
		level:  level,
	}
}

// Init sets up the global logger instance with the specified log level.
// It should be called from main before using the logger functions.
func Init(level LogLevel) {
	once.Do(func() {
		globalLogger = NewLogger(level)
	})
}

// logOutput is a helper function to format and output log messages.
func output(level LogLevel, levelStr string, format string, args ...interface{}) {
	if globalLogger.level <= level {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		} else {
			// shortfile := path.Base(path.Dir(file)) + "/" + path.Base(file)
			shortfile := path.Base(file)
			file = shortfile
		}
		timeStr := time.Now().Format("2006/01/02 15:04:05")
		color, ok := logLevelColors[level]
		if !ok {
			color = colorWhite // Default color
		}
		prefix := fmt.Sprintf("[%s]\t[%s]\t[%s:%d]\t", levelStr, timeStr, file, line)
		message := fmt.Sprintf(format, args...)
		globalLogger.Printf(color + prefix + message + colorReset)
	}
}

// Debug prints debug-level messages.
func Debug(format string, args ...interface{}) {
	output(DEBUG, "DEBUG", format, args...)
}

// Info prints information-level messages.
func Info(format string, args ...interface{}) {
	output(INFO, "INFO", format, args...)
}

// Warn prints warning-level messages.
func Warn(format string, args ...interface{}) {
	output(WARN, "WARN", format, args...)
}

// Error prints error-level messages.
func Error(format string, args ...interface{}) {
	output(ERROR, "ERROR", format, args...)
}

// Fatal prints fatal-level messages and exits the program.
func Fatal(format string, args ...interface{}) {
	output(FATAL, "FATAL", format, args...)
	os.Exit(1)
}
