package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received url with path %s, known urls: %v \n", r.URL, pathsToUrls)
		if url, ok := pathsToUrls[r.URL.Path]; ok {
			w.Header().Add("Location", url)
			w.WriteHeader(http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	data := make([]pathUrl, 0)
	if err := yaml.Unmarshal(yml, &data); err != nil {
		return nil, err
	}
	redirectsMap := make(map[string]string)
	fmt.Println("Deserialized yaml :", data)
	for i := range data {
		redirectsMap[data[i].Path] = data[i].URL
	}
	return MapHandler(redirectsMap, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
