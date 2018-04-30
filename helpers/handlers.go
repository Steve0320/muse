package helpers

import (
	"net/http"
	"log"
)

// Handle basic logging using native logger
func LoggingHandler(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("%s request to %s\n", r.Method, r.URL)
		//w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
