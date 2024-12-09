package pkg

import (
	"log"
	"net/http"
	"time"
)

type statusWriter struct {
	http.ResponseWriter
	status     int
	headerSent bool
}

func (w *statusWriter) WriteHeader(statusCode int) {
	if !w.headerSent {
		w.status = statusCode
		w.headerSent = true
		w.ResponseWriter.WriteHeader(statusCode)
	}
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sw, r)
		duration := time.Since(start)
		statusCode := sw.status
		log.Printf("%s %s %s %d %s %s", r.Method, r.URL.Path, r.Proto, statusCode, http.StatusText(statusCode), duration)
	})
}

func DisableCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))

		next.ServeHTTP(w, r)
	})
}
