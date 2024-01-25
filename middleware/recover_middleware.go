package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

type RecoverMiddleware struct {
	devMode bool
	handler http.Handler
}

func (m *RecoverMiddleware) ServeHTTP(r http.ResponseWriter, w *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			if m.devMode {
				r.Write([]byte(fmt.Sprintln(err)))
				stack := debug.Stack()
				r.Write(stack)
			} else {
				r.Write([]byte("Something went wrong"))
			}
		}
	}()
	m.handler.ServeHTTP(r, w)
}

func NewRecoverHandler(h http.Handler, devMode bool) http.Handler {
	return &RecoverMiddleware{devMode, h}
}
