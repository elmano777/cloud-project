package bootstrap

// @title API de Gestión de Cursos y Estudiantes
// @version 1.0
// @description API para la gestión de cursos y matrículas de estudiantes
// @termsOfService http://swagger.io/terms/
// @contact.name Soporte API
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"api-fiber/app/modules/cursos"
	estudiantecurso "api-fiber/app/modules/estudiante_cursos"
	"api-fiber/database/connections"

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	// Habilitar CORS para cualquier dominio
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		ExposeHeaders:    "Content-Length, Content-Type, Authorization",
		AllowCredentials: false,
		MaxAge:           86400, // Preflight requests can be cached for 24 hours
	}))

	app.Use(logger.New())
	app.Use(recover.New())

	app.Use(swagger.New(swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "API de Gestión de Cursos y Estudiantes",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	cursos.SetupCursoRoutes(app, queries)
	estudiantecurso.SetupEstudianteCursoRoutes(app, queries)

	fmt.Println("✅ Aplicación iniciada correctamente")
	return app
}
