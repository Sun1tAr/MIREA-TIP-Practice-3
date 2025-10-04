package api

import (
	"log"
	"net/http"
	"strings"
	"time"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.status = code
	sr.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: 200}
		next.ServeHTTP(rec, r)
		log.Printf("%s %s %d %v", r.Method, r.URL.Path, rec.status, time.Since(start))
	})
}

func CORSable(r *http.Request) bool {
	created := r.Header.Get("Access-Control-Allow-Origin") != ""
	validate := !strings.Contains(r.Header.Get("Access-Control-Allow-Origin"), "*")
	return created && validate
}

func TitleValidation(title string) bool {
	return len(title) >= 3 && len(title) < 140
}
