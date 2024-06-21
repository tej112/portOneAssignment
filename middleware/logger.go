package middleware

import (
	"log"
	"net/http"
	"time"
)

// Create a custom responseWriter to capture the status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Override the WriteHeader method to capture the status code
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Capture the current time
		now := time.Now()

		// Create a custom responseWriter
		rw := &responseWriter{ResponseWriter: w}

		// Call the next handler
		next.ServeHTTP(rw, r)

		// Calculate the elapsed time
		elapsed := time.Since(now)

		// Log the request details
		log.Println(r.Method, r.URL.Path, rw.statusCode, elapsed)
	})
}
