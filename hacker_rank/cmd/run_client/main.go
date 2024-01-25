package main

import (
	"fmt"
	"gophercises/hacker_rank/client"
	"gophercises/hacker_rank/stories"
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	c := client.NewHackerRankClient(client.WithDefaultApiBase())
	storiesService := stories.NewStories(c, stories.WithDefaultParams)
	mux := http.NewServeMux()
	pageTemplate := template.Must(template.New("page").Parse(page))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fetchedStories, err := storiesService.GetCachedStories()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Printf("Executed within %v \n", time.Now().Sub(start))
		pageTemplate.Execute(w, fetchedStories)
	})
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {

	})
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

const page = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Items List</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
        }

        h1 {
            color: #333;
        }

        ul {
            list-style-type: none;
            padding: 0;
        }

        li {
            border: 1px solid #ccc;
            padding: 10px;
            margin-bottom: 10px;
            border-radius: 5px;
        }

        a {
            color: #007bff;
            text-decoration: none;
        }
    </style>
</head>
<body>

<h1>Items List</h1>

<ul>
    {{range .}}
    <li>
        <strong>Title:</strong> {{.Title}}<br>
        <strong>By:</strong> {{.By}}<br>
        <strong>Score:</strong> {{.Score}}<br>
        <strong>Time:</strong> {{.Time}}<br>
        <strong>Type:</strong> {{.Type}}<br>
        {{if .Text}}
            <strong>Text:</strong> {{.Text}}<br>
        {{else}}
            <strong>URL:</strong> <a href="{{.URL}}" target="_blank">{{.URL}}</a><br>
        {{end}}
    </li>
    {{end}}
</ul>

</body>
</html>
`
