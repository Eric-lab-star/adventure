package main

import (
	"fmt"

	"log"
	"net/http"

	"os"

	"github.com/Eric-lab-star/adventure/pkg/story"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	json, err := os.Open("gopher.json")
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

	log.Print("listening to port 8080")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), story.NewHandler(data)))

}
