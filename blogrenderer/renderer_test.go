package blogrenderer_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	"github.com/gregor-pifko/learn-go-with-tests/blogrenderer"
	"github.com/gregor-pifko/learn-go-with-tests/reading-files/blogposts"
)

func TestRender(t *testing.T) {
	approvals.UseFolder("testdata")
	aPost := blogposts.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}

	renderPost := func(t testing.TB, post blogposts.Post) string {
		t.Helper()
		buf := bytes.Buffer{}
		if err := postRenderer.Render(&buf, post); err != nil {
			t.Fatal(err)
		}
		return buf.String()
	}

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		approvals.VerifyString(t, renderPost(t, aPost))
	})

	t.Run("it converts a single post with Markdown body to HTML", func(t *testing.T) {
		posts, _ := blogposts.NewPostFromFS(os.DirFS("./posts"))
		approvals.VerifyString(t, renderPost(t, posts[0]))
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogposts.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	aPost := blogposts.Post{
		Title:       "hello world",
		Body:        "This is a post",
		Description: "This is a description",
		Tags:        []string{"go", "tdd"},
	}

	postRenderer, err := blogrenderer.NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	for b.Loop() {
		postRenderer.Render(io.Discard, aPost)
	}
}
