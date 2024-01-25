package middleware

import (
	"fmt"
	"log"
	"net/http"
)

func DummyHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleHello)
	mux.HandleFunc("/panic", panicEndpoint)
	return mux
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte(fmt.Sprintf("Hello, %s \n", r.RemoteAddr)))
	if err != nil {
		log.Printf("Could not write in hello controller %s \n", r.RemoteAddr)
	}
}

func panicEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Written part of a panic response \n"))
	panic("Just a panic")
}
