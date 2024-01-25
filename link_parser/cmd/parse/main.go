package main

import (
	"flag"
	"fmt"
	"gophercises/link_parser"
	"log"
	"net/http"
	"os"
)

func main() {
	fileName := flag.String("f", "examples/ex3.html", "html fileName to parse")
	link := flag.String("link", "https://www.google.com", "link to parse html from")
	flag.Parse()
	file, err := os.Open(*fileName)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
		return
	}
	response, err := http.Get(*link)
	defer response.Body.Close()
	if err != nil {
		log.Fatalln(err)
		return
	}
	sl1, _ := link_parser.ExtractLinks(file)
	fmt.Println(sl1)
	sl2, _ := link_parser.ExtractLinks(response.Body)
	fmt.Println(sl2)
}
