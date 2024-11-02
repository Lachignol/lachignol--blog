+++
title = "Le code du blog"
description = "Main.go"
date = 2024-08-09

[author]
name = "La Chignol"
email = "Pas d'email for you "

[footer]
copyright = "©LaChignole"

+++



Voici le fichier `main.go` :

```go
package main

import (
	"bytes"
	"html/template"
	"strings"

	"io"
	"log"
	"net/http"
	"os"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

func main() {
	mux := http.NewServeMux()
	tpl, err := template.ParseFiles("posts.html")
	if err != nil {
		log.Fatal(err)
	}
	mux.HandleFunc("GET /post/{slug}", PostHandler(FileReader{}, tpl))

	err = http.ListenAndServe(":3030", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type SlugReader interface {
	Read(slug string) (string, error)
}

type FileReader struct{}

type Post struct {
	Title   string `toml:"title"`
	Slug    string `toml:"slug"`
	Content template.HTML
	Author  Author `toml:"author"`
	Footer  Footer `toml:"footer"`
}

type Author struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}
type Footer struct {
	Info string `toml:"info"`
}

func (fr FileReader) Read(slug string) (string, error) {
	file, err := os.Open(slug + ".md")
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil

}

func PostHandler(sl SlugReader, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		var post Post
		postMarkdown, err := sl.Read(slug)
		if err != nil {
			http.Error(w, "Post non trouvé", http.StatusNotFound)
			return

		}
		rest, err := frontmatter.Parse(strings.NewReader(postMarkdown), &post)
		if err != nil {
			http.Error(w, "Error parsing frontmatter", http.StatusInternalServerError)
			return
		}

		var buff bytes.Buffer
		mdStyle := goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(
					highlighting.WithStyle("dracula"),
				),
			))
		err = mdStyle.Convert([]byte(rest), &buff)
		if err != nil {
			http.Error(w, "Erreur lors de la conversion du markdown", http.StatusInternalServerError)
			return
		}
		post.Content = template.HTML(buff.String())
		err = tpl.Execute(w, post)
		if err != nil {
			http.Error(w, "Erreurr lors du parsing du template", http.StatusInternalServerError)
			return
		}
	}

}

```