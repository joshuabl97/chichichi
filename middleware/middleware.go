package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

func NewMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		writer := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(writer, r)

		latency := time.Since(start)

		statusCode := writer.Status()
		statusText := http.StatusText(statusCode)

		fmt.Printf("%s %s %s %s\n", r.Method, r.RequestURI, statusText, latency)
	})
}
