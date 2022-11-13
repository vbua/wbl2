package server

import (
	"log"
	"net/http"
	"time"
)

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%v, %v, %v, %s\n", r.Method, r.URL, timeStart.Format("2006-01-02T15:04:05"),
			time.Since(timeStart))
	})
}
