package cyoa

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

var t *template.Template

func init() {
	data, err := os.ReadFile("tmp_adventure.gohtml")
	if err != nil {
		return
	}
	t = template.Must(template.New("adventure").Parse(string(data)))
}

type AdventureHandler struct {
	adventure Adventure
}

func NewHandler(a Adventure) http.Handler {
	return AdventureHandler{a}
}

func (a AdventureHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := strings.Trim(r.URL.Path, "/")
	arc := a.adventure["intro"]
	if arcNext, ok := a.adventure[id]; ok {
		arc = arcNext
	}
	err := t.Execute(w, arc)
	if err != nil {
		log.Fatalln(err)
	}
}
