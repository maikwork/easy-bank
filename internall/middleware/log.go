package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func LoggerMiddlerware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
