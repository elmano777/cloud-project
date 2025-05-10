package cursos

import (
	"api-fiber/database/generated"

	"github.com/gofiber/fiber/v2"
)

func SetupCursoRoutes(app *fiber.App, queries *generated.Queries) {
	cursoRepo := NewCursoRepository(queries)
	cursoService := NewCursoService(cursoRepo)
	cursoController := NewCursoController(cursoService)

	app.Get("/cursos/test", cursoController.Testing)
	app.Post("/cursos", cursoController.CreateCurso)
	app.Get("/cursos/:codigo", cursoController.GetCursoByCodigo)
	app.Get("/cursos", cursoController.ListCursos)
	app.Put("/cursos/:codigo", cursoController.UpdateCurso)
	app.Delete("/cursos/:codigo", cursoController.DeleteCurso)
}
