package main

import (
	"flag"
	"gophercises/sitemap_builder"
)

func main() {
	url := flag.String("url", "https://www.calhoun.io/", "Provide url to parse")
	depth := flag.Int("d", 1, "Provide depth to parse to parse")
	flag.Parse()
	sitemap_builder.TraverseUrl(*url, *depth)
}
