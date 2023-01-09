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

func NewHandler(story Story) http.Handler {
	return handler{story}
}

type handler struct {
	story Story
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("../static/template.gohtml"))
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, " ")
	if path == "/" || path == "" {
		path = "/intro"
	}
	chapter, ok := h.story[path[1:]]
	if ok {
		err := tmpl.Execute(w, chapter)
		if err != nil {
			fmt.Printf("tmpl.execute error:\n%v\n", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
	http.Error(w, "Page not found", http.StatusNotFound)

}
