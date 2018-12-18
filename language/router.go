package language

import (
	"github.com/go-chi/chi"
	"net/http"
)

func LanguageRouter() http.Handler {
	r := chi.NewRouter()
	r.Post("/", createLanguage)
	r.Get("/{languageID}", getLanguage)
	r.Get("/list", getLanguageList)
	r.Delete("/{languageID}", removeLanguage)
	r.Put("/{languageID}", updateLanguage)
	return r
}
