package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/mdn":  "https://developer.mozilla.org",
		"/leet": "https://leetcode.com/mishamolnar/",
	}
	mapHandler := MapHandler(pathsToUrls, mux)
	filename := flag.String("f", "redirects.yaml", "Specify filename of csv files with problems")
	yml, err := os.ReadFile(*filename)
	if err != nil {
		log.Println(err)
	}
	yamlHandler, err := YAMLHandler(yml, mapHandler)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Listening on localhost:8080")
	http.ListenAndServe("localhost:8080", yamlHandler)
}

// default http request multiplexer with function for the root
// mux is also Handler from net/http
func defaultMux() *http.ServeMux {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello from root")
	if err != nil {
		log.Println(err)
	}
}
