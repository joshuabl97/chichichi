package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type statusLoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusLoggingResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func NewMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		writer := &statusLoggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(writer, r)

		latency := time.Since(start)

		statusCode := writer.statusCode
		statusText := http.StatusText(statusCode)

		fmt.Printf("%s %s %s %s\n", r.Method, r.RequestURI, statusText, latency)
	})
}

func Logger(handler http.Handler) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/", handler)
	return r
}
