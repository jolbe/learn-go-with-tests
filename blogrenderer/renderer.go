// Package blogrenderer is responsible for rendering posts as HTML
package blogrenderer

import (
	"bytes"
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gregor-pifko/learn-go-with-tests/reading-files/blogposts"
	"github.com/yuin/goldmark"
)

//go:embed "templates/*"
var postTemplates embed.FS

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "**/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ}, nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []blogposts.Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", mapToPostVMs(posts))
}

func (r *PostRenderer) Render(w io.Writer, p blogposts.Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p))
}

type postViewModel struct {
	Title, SanitizeTitle, Description, Body string
	BodyHTML                                template.HTML
	Tags                                    []string
}

func newPostVM(p blogposts.Post) postViewModel {
	bodyHTML, _ := convertToHTML(p.Body) // ignore conversion errors to keep a single bad post from breaking the entire index
	return postViewModel{
		Title:         p.Title,
		SanitizeTitle: strings.ToLower(strings.ReplaceAll(p.Title, " ", "-")),
		Description:   p.Description,
		Body:          p.Body,
		BodyHTML:      template.HTML(bodyHTML),
		Tags:          p.Tags,
	}
}

func convertToHTML(markdown string) (string, error) {
	buf := bytes.Buffer{}
	if err := goldmark.Convert([]byte(markdown), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func mapToPostVMs(posts []blogposts.Post) []postViewModel {
	var postViews []postViewModel
	for _, p := range posts {
		postViews = append(postViews, newPostVM(p))
	}
	return postViews
}
