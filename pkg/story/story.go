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

type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.tmpl = t
	}
}

func WithPathFn(pathFn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.path = pathFn
	}
}

func NewHandler(story Story, opts ...HandlerOption) http.Handler {
	h := handler{story, defaultTmpl, defaultPath}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	story Story
	tmpl  *template.Template
	path  func(r *http.Request) string
}

func init() {

	defaultTmpl = template.Must(template.ParseFiles("public/template.gohtml"))
}

func defaultPath(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/" || path == "" {
		path = "/intro"
	}
	return path
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.path(r)
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
