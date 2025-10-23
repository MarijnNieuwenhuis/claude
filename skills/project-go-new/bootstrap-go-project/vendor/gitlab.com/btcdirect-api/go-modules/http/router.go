package http

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Returns a new router with logging middleware.
func createLoggingRouter(r *mux.Router, log *zap.SugaredLogger) http.Handler {
	return loggingRouter(r, log)
}

// Override ResponseWriter to inject HTTP status code.
func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Logging middleware for HTTP requests.
// This middleware logs the HTTP request and its response status code.
//
// The log message will be formatted as follows:
//
// <host> - <method> <path> - <status code> <protocol>
//
// Example:
//
// 8.8.8.8 - GET /health - 200 HTTP/1.1
func loggingRouter(handler http.Handler, log *zap.SugaredLogger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := &loggingResponseWriter{w, http.StatusOK}
		handler.ServeHTTP(lrw, r)

		statusCode := lrw.statusCode
		host, _, err := net.SplitHostPort(r.RemoteAddr)

		if err != nil {
			host = r.RemoteAddr
		}

		// Log the HTTP request
		log.Infof("%s - %s %s - %d %s", host, r.Method, r.URL.Path, statusCode, r.Proto)
	})
}
