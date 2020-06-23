package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type EventsHandler struct {
	log logrus.FieldLogger
}

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewEventsHandler(log logrus.FieldLogger) *EventsHandler {
	return &EventsHandler{
		log: log,
	}
}

func (e *EventsHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{
		"status": "UP",
	})
	if err != nil {
		e.log.Printf("error in sending response : %w", err)
	}
}

func (e *EventsHandler) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapperWriter := NewLoggingResponseWriter(w)
		next.ServeHTTP(wrapperWriter, r)
		latency := time.Since(start)
		e.log.Printf("%s %s %s %s %dms %d %s",
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			r.Proto,
			latency.Microseconds(),
			wrapperWriter.statusCode,
			r.UserAgent())
	})
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
