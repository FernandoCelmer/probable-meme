package main

import (
	web "github.com/FernandoCelmer/api.go/src"
	"github.com/FernandoCelmer/probable-meme/handlers"
)

func main() {
	app := web.NewApp(
		web.Title("API"),
		web.Description("Probable Meme"),
		web.Version("0.1.0"),
	)

	app.Get("/item", handlers.ItemHandler)

	app.Run(web.Port(8080))

}
