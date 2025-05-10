package estudiantecurso

import (
	"api-fiber/database/generated"

	"github.com/gofiber/fiber/v2"
)

func SetupEstudianteCursoRoutes(app *fiber.App, queries *generated.Queries) {
	estudianteCursoRepo := NewEstudianteCursoRepository(queries)

	userClient := NewUserClient()

	estudianteCursoService := NewEstudianteCursoService(estudianteCursoRepo, userClient)

	estudianteCursoController := NewEstudianteCursoController(estudianteCursoService)

	app.Get("/estudiante-cursos/test", estudianteCursoController.Testing)
	app.Post("/estudiante-cursos", estudianteCursoController.Inscribir)
	app.Delete("/estudiante-cursos", estudianteCursoController.Desinscribir)
	app.Get("/estudiante-cursos/estudiante/:id", estudianteCursoController.GetCursosByEstudiante)
	app.Get("/estudiante-cursos/curso/:codigo", estudianteCursoController.GetEstudiantesByCurso)
}
