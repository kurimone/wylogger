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

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

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

var logLevelColors = map[LogLevel]string{
	DEBUG: colorBlue,
	INFO:  colorGreen,
	WARN:  colorYellow,
	ERROR: colorRed,
	FATAL: colorPurple,
}

type Logger struct {
	*log.Logger
	level LogLevel
}

var (
	globalLogger *Logger
	once         sync.Once
)

func NewLogger(level LogLevel) *Logger {
	return &Logger{
		Logger: log.New(os.Stdout, "", 0), // Disable the default timestamp
		level:  level,
	}
}

func Init(level LogLevel) {
	once.Do(func() {
		globalLogger = NewLogger(level)
	})
}

func SetLevel(level LogLevel) {
	if globalLogger != nil {
		globalLogger.level = level
	}
}

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

func Debug(format string, args ...interface{}) {
	output(DEBUG, "DEBUG", format, args...)
}

func Info(format string, args ...interface{}) {
	output(INFO, "INFO", format, args...)
}

func Warn(format string, args ...interface{}) {
	output(WARN, "WARN", format, args...)
}

func Error(format string, args ...interface{}) {
	output(ERROR, "ERROR", format, args...)
}

func Fatal(format string, args ...interface{}) {
	output(FATAL, "FATAL", format, args...)
	os.Exit(1)
}
