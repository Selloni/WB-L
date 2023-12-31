package handler

import (
	"log"
	"net/http"
	"time"
)

// записываем в логи данные при запросе
func middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s request: %s, %s\n", time.Now(), r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}
