package estudiantecurso

import (
	"api-fiber/database/generated"

	"github.com/gofiber/fiber/v2"
)

func SetupEstudianteCursoRoutes(app *fiber.App, queries *generated.Queries) {
	// Repositorio
	estudianteCursoRepo := NewEstudianteCursoRepository(queries)

	// Cliente para validar el estudiante
	userClient := NewUserClient()

	// Servicio
	estudianteCursoService := NewEstudianteCursoService(estudianteCursoRepo, userClient)

	// Controlador
	estudianteCursoController := NewEstudianteCursoController(estudianteCursoService)

	// Rutas de estudiante-cursos
	app.Post("/estudiante-cursos", estudianteCursoController.Inscribir)                           // POST /estudiante-cursos
	app.Delete("/estudiante-cursos", estudianteCursoController.Desinscribir)                      // DELETE /estudiante-cursos
	app.Get("/estudiante-cursos/estudiante/:id", estudianteCursoController.GetCursosByEstudiante) // GET /estudiante-cursos/estudiante/:id
	app.Get("/estudiante-cursos/curso/:codigo", estudianteCursoController.GetEstudiantesByCurso)  // GET /estudiante-cursos/curso/:codigo
}
