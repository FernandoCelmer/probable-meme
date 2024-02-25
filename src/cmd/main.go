package main

import (
	"github.com/FernandoCelmer/probable-meme/src/handlers"

	web "github.com/FernandoCelmer/api.go/src"
)

func main() {
	app := web.NewApp(
		web.Title("API"),
		web.Description("Probable Meme"),
		web.Version("0.1.0"),
	)

	app.Post("/clientes/{id}/transacoes", handlers.ClientTransacoesHandler)
	app.Get("/clientes/{id}/extrato", handlers.ClientesExtratoHandler)

	app.Run(web.Port(8080))

}
