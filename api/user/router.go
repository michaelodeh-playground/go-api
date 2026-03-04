package user

import (
	"github.com/go-chi/chi/v5"
)

var UserRoute = chi.NewRouter()

func init() {

	UserRoute.Post("/", CreateUser)
	UserRoute.Get("/", IndexUser)
	UserRoute.Put("/{user}", UpdateUser)
}
