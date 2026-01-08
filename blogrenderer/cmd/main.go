package main

import (
	"os"

	"github.com/gregor-pifko/learn-go-with-tests/blogrenderer"
	"github.com/gregor-pifko/learn-go-with-tests/reading-files/blogposts"
)

func main() {
	posts, _ := blogposts.NewPostFromFS(os.DirFS("./posts"))
	renderer, _ := blogrenderer.NewPostRenderer()
	renderer.Render(os.Stdout, posts[0])
}
