package main

import (
	"log"
	"os"

	"github.com/gregor-pifko/learn-go-with-tests/reading-files/blogposts"
)

func main() {
	posts, err := blogposts.NewPostFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
