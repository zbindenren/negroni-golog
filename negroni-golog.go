package negronigolog

import (
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	"github.com/timehop/golog/log"
)

// Logger is a middleware handler that logs the request as it goes in and the response as it goes out.
type Logger struct {
	// Logger is the log.Logger instance used to log messages with the Logger middleware
	*log.Logger
}

// NewLogger returns the standard golog logger.
func NewLogger() *Logger {
	logger := log.New()
	flags := log.FlagsDate | log.FlagsTime
	logger.SetTimestampFlags(flags)
	return &Logger{logger}
}

// NewLoggerWithPrefix returns a logger with a custom prefix.
func NewLoggerWithPrefix(prefix string) *Logger {
	logger := NewLogger()
	logger.ID = prefix
	return logger
}

// NewLoggerWithPrefixAndFlags returns a logger with a custom prefix and timestamp flags.
func NewLoggerWithPrefixAndFlags(prefix string, flags int) *Logger {
	logger := log.New()
	logger.ID = prefix
	logger.SetTimestampFlags(flags)
	return &Logger{logger}
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	l.Info("start", "url", r.URL.Path, "method", r.Method)

	next(rw, r)

	res := rw.(negroni.ResponseWriter)
	l.Info("completed", "url", r.URL.Path, "status", res.Status(), "status_text", http.StatusText(res.Status()), "took", time.Since(start))
}
