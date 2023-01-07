package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"os"

	"github.com/Eric-lab-star/adventure/pkg/story"
)

func main() {

	r, err := os.Open("../gopher.json")
	if err != nil {
		fmt.Print("failed to open file\n")
		os.Exit(1)
	}
	story, err := story.JsontoStory(r)
	if err != nil {
		fmt.Println("failed to parse json to story")
		os.Exit(1)
	}
	for k, v := range story {
		pattern := fmt.Sprint("/", k)
		http.HandleFunc(pattern, chapter(v))
	}

	log.SetFlags(log.Ltime)
	log.Print("listening to http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}

func chapter(data story.Chapter) http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("../static/template.gohtml"))
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, data)
		if err != nil {
			fmt.Println("err: tmpl execute")
			os.Exit(1)
		}

	}
}
