package story

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../static/template.gohtml"))
	err := tmpl.Execute(w, h.story["intro"])
	if err != nil {
		fmt.Println("exetue err")
		os.Exit(1)
	}
}
