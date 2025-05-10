package main

import (
	"log"

	bootstrap "api-fiber/app/boostrap"
	_ "api-fiber/docs"
)

// @title API de Gestión de Cursos y Estudiantes
// @version 1.0
// @description API para la gestión de cursos y matrículas de estudiantes
// @termsOfService http://swagger.io/terms/
// @contact.name Soporte API
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8070
// @BasePath /
func main() {
	app := bootstrap.InitializeApp()

	log.Println("🚀 Servidor iniciado en http://localhost:8070")
	app.Listen(":8070")
}
