package global

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

type MetadataRecup interface {
	Query() ([]PostMetadata, error)
}

type PostMetadata struct {
	Slug        string
	Title       string    `toml:"title"`
	Author      Author    `toml:"author"`
	Description string    `toml:"description"`
	Date        time.Time `toml:"date"`
	Footer      Footer    `toml:"footer"`
}

type SlugReader interface {
	Read(slug string) (string, error)
}

type FileReader struct {
	Dir string
}

type PostData struct {
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
	Copyright string `toml:"copyright"`
}

func (fr FileReader) Read(slug string) (string, error) {
	slugPath := filepath.Join(fr.Dir, slug+".md")
	file, err := os.Open(slugPath)
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

func (fr FileReader) Query() ([]PostMetadata, error) {
	postPath := filepath.Join(fr.Dir, "*.md")
	filenames, err := filepath.Glob(postPath)
	if err != nil {
		return nil, fmt.Errorf("Requete de recuperationde fichiers: %w", err)
	}
	var posts []PostMetadata
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return nil, fmt.Errorf("Ouverture de fichier %s :%w", filename, err)
		}
		defer f.Close()
		var post PostMetadata
		_, err = frontmatter.Parse(f, &post)
		if err != nil {
			return nil, fmt.Errorf("Recuperation des metadata pour le fichier %s:%w", filename, err)
		}
		post.Slug = strings.TrimSuffix(filepath.Base(filename), ".md")

		posts = append(posts, post)

	}
	return posts, nil
}

func PostHandler(sl SlugReader, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug")
		postMarkdown, err := sl.Read(slug)
		if err != nil {
			http.Error(w, "Post non trouvé", http.StatusNotFound)
			return

		}
		var post PostData
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
			http.Error(w, "Erreur lors du parsing du template", http.StatusInternalServerError)
			return
		}
	}
}

type IndexData struct {
	Posts []PostMetadata
}

func IndexHandler(mq MetadataRecup, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := mq.Query()
		if err != nil {
			http.Error(w, "Erreur lors de la recuperation des posts", http.StatusInternalServerError)
			return
		}
		data := IndexData{
			Posts: posts,
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Erreur lors de l'execution du template", http.StatusInternalServerError)
			return
		}
	}
}
