package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	global "github.com/Lachignol/lachignol-blog"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)

	}
}

func run(args []string, stdout io.Writer) error {
	mux := http.NewServeMux()

	postreader := global.FileReader{
		Dir: "posts",
	}
	postTemplate := template.Must(template.ParseFiles("posts.html"))
	indexTemplate := template.Must(template.ParseFiles("index.html"))
	mux.HandleFunc("GET /post/{slug}", global.PostHandler(postreader, postTemplate))
	mux.HandleFunc("GET /", global.IndexHandler(postreader, indexTemplate))
	err := http.ListenAndServe(":3030", mux)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
