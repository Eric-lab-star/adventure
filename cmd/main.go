package main

import (
	"fmt"
	"text/template"

	"log"
	"net/http"

	"os"

	"github.com/Eric-lab-star/adventure/pkg/story"
)

func main() {

	json, err := os.Open("../gopher.json")
	if err != nil {
		fmt.Print("failed to open file\n")
		os.Exit(1)
	}
	data, err := story.JsonDecoder(json)
	if err != nil {
		fmt.Println("failed to parse json to story")
		os.Exit(1)
	}

	log.SetFlags(log.Ltime)
	temp := template.Must(template.New("").Parse("hello"))
	log.Print("listening to http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", story.NewHandler(data, story.WithTemplate(temp))))

}
