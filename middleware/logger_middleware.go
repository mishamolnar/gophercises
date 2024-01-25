package middleware

import (
	"log"
	"net/http"
	"time"
)

type LoggerHandler struct {
	handler http.Handler
}

func (lh *LoggerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	lh.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func NewLoggerMiddleware(h http.Handler) *LoggerHandler {
	return &LoggerHandler{h}
}
