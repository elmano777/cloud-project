package main

import (
	"log"

	bootstrap "api-fiber/app/boostrap"
)

func main() {
	app := bootstrap.InitializeApp()

	log.Println("🚀 Servidor iniciado en http://localhost:8070")
	app.Listen(":8070")
}
