package api

import (
	"api/api/user"
	"api/api/wallet"

	"github.com/go-chi/chi/v5"
)

var ApiRoute = chi.NewRouter()

func init() {
	ApiRoute.Mount("/users", user.Routes)
	ApiRoute.Mount("/wallets", wallet.Routes)
}
