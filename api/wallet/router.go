package wallet

import "github.com/go-chi/chi/v5"

var Routes = chi.NewRouter()

func init() {
	Routes.Post("/fund", Found)
	Routes.Get("/balance/{user}", Balance)
}
