package story

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"
)

func JsonDecoder(r io.Reader) (Story, error) {
	story := Story{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&story)

	if err != nil {

		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

var defaultTmpl *template.Template

/*
if t is not provided default template will be used
*/
func NewHandler(story Story, t *template.Template) http.Handler {
	if t == nil {
		t = defaultTmpl
	}
	return handler{story, t}
}

type handler struct {
	story Story
	tmpl  *template.Template
}

func init() {
	defaultTmpl = template.Must(template.ParseFiles("../static/template.gohtml"))
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, " ")
	if path == "/" || path == "" {
		path = "/intro"
	}

	if chapter, ok := h.story[path[1:]]; ok {
		err := h.tmpl.Execute(w, chapter)
		if err != nil {
			fmt.Printf("tmpl.execute error:\n%v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)

		}
		return
	}

	http.Error(w, "Page not found", http.StatusNotFound)

}
