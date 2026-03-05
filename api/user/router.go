package user

import (
	"github.com/go-chi/chi/v5"
)

var Routes = chi.NewRouter()

func init() {
	Routes.Get("/", Index)
	Routes.Get("/{user}", Show)
	Routes.Post("/", Create)
	Routes.Put("/{user}", Update)
	Routes.Delete("/{user}", Delete)
}
