package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/pkg/models"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/services/calendar"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

type CalendarAPIServer struct {
	log             logrus.FieldLogger
	calendarService *calendar.Calendar
	host            string
	server          *http.Server
}

func NewCalendarAPIServer(log logrus.FieldLogger, host string, calendarService *calendar.Calendar) *CalendarAPIServer {
	service := &CalendarAPIServer{
		log:             log,
		calendarService: calendarService,
		host:            host,
	}
	router := mux.NewRouter()
	router.Use(service.ContentTypeMiddleware)
	router.Use(service.LoggingMiddleware)

	router.HandleFunc("/health", service.HealthCheck).Methods("GET")

	// events edpoint
	router.HandleFunc("/events", service.CreateEvent).Methods("PUT")
	router.HandleFunc("/events/{id}", service.UpdateEvent).Methods("POST")
	router.HandleFunc("/events/{id}", service.DeleteEventByID).Methods("DELETE")
	router.HandleFunc("/events", service.GetAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", service.FindEventByID).Methods("GET")
	router.HandleFunc("/events/day", service.FindDayEvents).Methods("GET")
	router.HandleFunc("/events/week", service.FindWeekEvents).Methods("GET")
	router.HandleFunc("/events/month", service.FindMonthEvents).Methods("GET")

	server := http.Server{
		Addr:         host,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	service.server = &server
	return service
}

func (c *CalendarAPIServer) ListenAndServe() error {
	notifyCh := make(chan os.Signal, 1)
	errorCh := make(chan error)
	signal.Notify(notifyCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-notifyCh
		c.log.Println("Stopping server")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		c.server.SetKeepAlivesEnabled(false)
		if err := c.server.Shutdown(ctx); err != nil {
			errorCh <- err
		}
		close(errorCh)
	}()

	c.log.Printf("Start server on %s....", c.host)
	if err := c.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	if err := <-errorCh; err != nil {
		return err
	}
	c.log.Println("Stop server successfully")
	return nil
}

func (c *CalendarAPIServer) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapperWriter := NewLoggingResponseWriter(w)
		next.ServeHTTP(wrapperWriter, r)
		latency := time.Since(start)
		c.log.Printf("%s %s %s %s %dms %d %s",
			r.RemoteAddr,
			r.Method,
			r.RequestURI,
			r.Proto,
			latency.Microseconds(),
			wrapperWriter.statusCode,
			r.UserAgent())
	})
}

func (c *CalendarAPIServer) ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (c *CalendarAPIServer) HealthCheck(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(map[string]string{
		"status": "UP",
	})
	if err != nil {
		c.log.Printf("error in sending response : %w", err)
	}
}

func (c *CalendarAPIServer) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event := &models.Event{}
	err := json.NewDecoder(r.Body).Decode(event)
	if err != nil {
		c.log.Printf("failed to get body %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := c.calendarService.CreateEvent(context.Background(), event)
	if err != nil {
		c.log.Printf("failed to create event %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(map[string]string{
		"id": id.String(),
	})

	if err != nil {
		c.log.Printf("error in sending response : %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *CalendarAPIServer) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	event := &models.Event{}
	err := json.NewDecoder(r.Body).Decode(event)
	if err != nil {
		c.log.Printf("failed to get body %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rID := mux.Vars(r)["id"]
	eID, err := uuid.Parse(rID)
	if err != nil {
		c.log.Printf("failed to parse id %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.calendarService.UpdateEvent(context.Background(), eID, event)
	if err != nil {
		c.log.Printf("failed to update event %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *CalendarAPIServer) DeleteEventByID(w http.ResponseWriter, r *http.Request) {
	rID := mux.Vars(r)["id"]
	eID, err := uuid.Parse(rID)
	if err != nil {
		c.log.Printf("failed to parse id %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.calendarService.DeleteEvent(context.Background(), eID)
	if err != nil {
		c.log.Printf("failed to delete event %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *CalendarAPIServer) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := c.calendarService.GetAllEvents(context.Background())
	if err != nil {
		c.log.Printf("failed to update event %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		c.log.Printf("failed to get all events %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *CalendarAPIServer) FindEventByID(w http.ResponseWriter, r *http.Request) {
	rID := mux.Vars(r)["id"]
	eID, err := uuid.Parse(rID)
	if err != nil {
		c.log.Printf("failed to parse id %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	event, err := c.calendarService.GetEventByID(context.Background(), eID)
	if err != nil {
		c.log.Printf("failed to find event %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(event)
	if err != nil {
		c.log.Printf("failed write response %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *CalendarAPIServer) FindDayEvents(w http.ResponseWriter, r *http.Request) {
	sDay := r.FormValue("from")
	startDay, err := time.Parse("2006-01-02T15:04:05-0700", sDay)
	if err != nil {
		c.log.Printf("failed to parse from day %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := c.calendarService.FindDayEvents(context.Background(), startDay)
	if err != nil {
		c.log.Printf("failed to find day events %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		c.log.Printf("failed to write response %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *CalendarAPIServer) FindWeekEvents(w http.ResponseWriter, r *http.Request) {
	sDay := r.FormValue("from")
	startDay, err := time.Parse("2006-01-02T15:04:05-0700", sDay)
	if err != nil {
		c.log.Printf("failed to parse from day %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := c.calendarService.FindWeekEvents(context.Background(), startDay)
	if err != nil {
		c.log.Printf("failed to find week events %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		c.log.Printf("failed to write response %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *CalendarAPIServer) FindMonthEvents(w http.ResponseWriter, r *http.Request) {
	sDay := r.FormValue("from")
	startDay, err := time.Parse("2006-01-02T15:04:05-0700", sDay)
	if err != nil {
		c.log.Printf("failed to parse from day %w", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	events, err := c.calendarService.FindWeekEvents(context.Background(), startDay)
	if err != nil {
		c.log.Printf("failed to find month events %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&events)
	if err != nil {
		c.log.Printf("failed to write response %w", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
