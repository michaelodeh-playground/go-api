package main

import (
	"api/app"
	"api/docs"
	_ "api/docs"
)

func main() {
	docs.SwaggerInfo.Title = "API Tutorial"
	app.Boot()
	app.Start()
}
