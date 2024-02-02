package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

var urls = map[string]string{
	"mdn":   "https://developer.mozilla.org/en-US/",
	"react": "https://react.dev/",
	"wiki":  "https://uk.wikipedia.org/wiki",
}

func main() {
	mux := http.ServeMux{}
	mux.HandleFunc("/", handleSearch)
	log.Fatal(http.ListenAndServe("localhost:8080", &mux))
}

// parse timeout from request parameters, when present - create context with it, when no -> then just context
func handleSearch(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	timeout, err := time.ParseDuration(r.FormValue("t"))
	if err != nil {
		ctx, cancel = context.WithCancel(context.Background())
	} else {
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	}
	defer cancel()
	handleQuery(ctx, w, r)
}

// retrieves q from FormValue, redirects to another page if the value was under timeout, otherwise not
func handleQuery(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	key := "test"
	if len(r.URL.Query()["q"]) > 0 {
		key = r.URL.Query()["q"][0]
	}
	fmt.Println(key)
	url, ok := urls[key]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ch := make(chan []byte)
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			ch <- []byte("Could not fetch desired page")
			return
		}
		defer resp.Body.Close()
		buf := new(strings.Builder)
		_, err = io.Copy(buf, resp.Body)
		if err != nil {
			ch <- []byte("Could not get body")
			return
		}
		ch <- []byte(buf.String())
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-ch:
			w.Write(data)
		}
	}
}
