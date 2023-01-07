package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	r, err := os.Open("gopher.json")
	if err != nil {
		fmt.Print("failed to open file")
		os.Exit(1)
	}
	d := json.NewDecoder(r)
	story := Story{}
	err = d.Decode(&story)
	if err != nil {
		fmt.Print("failed to decode json file")
		os.Exit(1)
	}
	fmt.Printf("%+v\n", story)

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
