// main.go
package main

import (
	bootstrap "api-fiber/app/boostrap"
	_ "api-fiber/docs"
	"log"
)

func main() {
	app := bootstrap.InitializeApp()
	log.Println("🚀 Servidor iniciado en http://localhost:8070")
	app.Listen(":8070")
}
