package webserver

import (
	"embed"
	"html/template"
	"net/http"

	"github.com/gpifko/go-specs-greet/domain/interactions"
)

const (
	greetPath = "/greet"
	cursePath = "/curse"
)

//go:embed "markup/*"
var templates embed.FS

func NewHandler() (http.Handler, error) {
	templ, err := template.ParseFS(templates, "markup/*.gohtml")
	if err != nil {
		return nil, err
	}

	handler := handler{templ: templ}
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.form)
	mux.HandleFunc(greetPath, handler.replyWith(interactions.Greet))
	mux.HandleFunc(cursePath, handler.replyWith(interactions.Curse))
	return mux, nil
}

type handler struct {
	templ *template.Template
}

func (h handler) replyWith(interact func(string) string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := h.templ.ExecuteTemplate(w, "reply.gohtml", interact(r.FormValue("name"))); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h handler) form(w http.ResponseWriter, r *http.Request) {
	_ = h.templ.ExecuteTemplate(w, "form.gohtml", nil)
}
