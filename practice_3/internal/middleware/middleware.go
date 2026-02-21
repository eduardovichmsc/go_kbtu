package middleware

import (
	"log"
	"net/http"
	"time"
)

const ValidAPIKey = "secret-key-123"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey == "" || apiKey != ValidAPIKey {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"error": "401-Unauthorized"}`))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		next.ServeHTTP(w, r)

		log.Printf("[TIMESTAMP]: %s | [HTTP METHOD]: %s | [ENDPOINT NAME]: %s",
			time.Now().Format(time.RFC3339),
			r.Method,
			r.URL.Path,
		)
		log.Printf("Request took: %v", time.Since(start))
	})
}