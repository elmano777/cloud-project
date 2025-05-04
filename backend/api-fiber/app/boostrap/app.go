package bootstrap

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"api-fiber/app/modules/cursos"
	estudiantecurso "api-fiber/app/modules/estudiante_cursos"
	"api-fiber/database/connections"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func InitializeApp() *fiber.App {
	dbPool, queries, err := connections.InitDB()
	if err != nil {
		log.Fatalf("Error conectando a la base de datos: %v", err)
	}

	cleanup := func() {
		dbPool.Close()
		log.Println("✅ Pool de conexiones a la base de datos cerrado correctamente")
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		cleanup()
		os.Exit(0)
	}()

	err = godotenv.Load()
	if err != nil {
		log.Println("⚠️ No se pudo cargar el archivo .env, usando variables de entorno del sistema.")
	}

	appEnv := os.Getenv("APP_ENV")
	log.Println("📌 APP_ENV:", appEnv)

	app := fiber.New(fiber.Config{
		AppName: "Aplicación de Cursos y Estudiantes",
	})

	app.Use(logger.New())
	app.Use(recover.New())

	// Rutas básicas
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Rutas de autenticación y perfil
	cursos.SetupCursoRoutes(app, queries)
	estudiantecurso.SetupEstudianteCursoRoutes(app, queries) // Aquí se configuran las rutas de estudiante-cursos

	fmt.Println("✅ Aplicación iniciada correctamente")
	return app
}
