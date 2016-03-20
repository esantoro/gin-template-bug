package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	t, err := template.ParseFiles("layout.html", "index.html")
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Execute(os.Stdout, nil); err != nil {
		log.Fatal(err)
	}

}
