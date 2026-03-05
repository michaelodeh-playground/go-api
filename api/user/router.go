package user

import (
	"github.com/go-chi/chi/v5"
)

var UserRoute = chi.NewRouter()

func init() {
	UserRoute.Get("/", Index)
	UserRoute.Get("/{user}", Show)
	UserRoute.Post("/", Create)
	UserRoute.Put("/{user}", Update)
	UserRoute.Delete("/{user}", Delete)
}
