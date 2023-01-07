package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.ReadFile("gopher.json")
	if err != nil {
		log.Fatal(err)
	}
	data := Story{}
	err = json.Unmarshal(f, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", data["intro"])

}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"chapter"`
}
