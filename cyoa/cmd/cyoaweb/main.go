package main

import (
	"flag"
	"fmt"
	"gophercises/cyoa"
	"log"
	"net/http"
	"os"
)

func main() {
	filename := flag.String("f", "data/gopher.json", "filename with adventure")
	flag.Parse()
	fmt.Printf("Using file %s for adventure parsing \n", *filename)
	data, err := os.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}
	a, err := cyoa.GetAdventure(data)
	if err != nil {
		log.Fatal(err)
	}
	mainHandler := cyoa.NewHandler(a)
	err = http.ListenAndServe("localhost:8080", mainHandler)
}
