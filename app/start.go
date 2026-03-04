package app

import (
	"fmt"
	"net/http"
	"os"
)

func Start() {
	r := Routes()
	port := os.Getenv("PORT")
	fmt.Println("Server running on port", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
