package api

import (
	"api/api/user"

	"github.com/go-chi/chi/v5"
)

var ApiRoute = chi.NewRouter()

func init() {
	ApiRoute.Mount("/users", user.UserRoute)
}
