package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gpifko/go-specs-greet/domain/interactions"
)

const (
	greetPath = "/greet"
	cursePath = "/curse"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc(greetPath, replyWith(interactions.Greet))
	mux.HandleFunc(cursePath, replyWith(interactions.Curse))
	return mux
}

func replyWith(f func(string) string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Fprint(w, f(name)) //nolint:errcheck
	}
}
