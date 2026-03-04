package app

import (
	"api/api"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", index)
	r.Mount("/api", api.ApiRoute)
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	return r
}
