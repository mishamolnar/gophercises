package main

import (
	"gophercises/middleware"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe("localhost:8080", middleware.NewLoggerMiddleware(middleware.NewRecoverHandler(middleware.DummyHandler(), true))))
}
