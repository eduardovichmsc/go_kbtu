package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const AUTHHEADER = "X-API-KEY"
const SECRETKEY = "spirited"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			key := r.Header.Get(AUTHHEADER)

			if key != SECRETKEY {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized);
				json.NewEncoder(w).Encode(map[string]string{"error" : "Unauthorized"})
				return;
			}

			next.ServeHTTP(w, r)
		})
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("%s %s %s", start.Format("2001-09-11T12:34:45"), r.Method, r.URL.Path)
	})
}
