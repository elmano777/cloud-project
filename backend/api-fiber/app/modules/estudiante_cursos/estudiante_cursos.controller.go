package estudiantecurso

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type EstudianteCursoController struct {
	service *EstudianteCursoService
}

func NewEstudianteCursoController(service *EstudianteCursoService) *EstudianteCursoController {
	return &EstudianteCursoController{service: service}
}

// Testing godoc
// @Summary Endpoint de prueba para estudiante-cursos
// @Description Verifica si el controlador de estudiante-cursos está funcionando correctamente
// @Tags estudiante-cursos
// @Produce plain
// @Success 200 {string} string "Endpoint de prueba para estudiante-cursos"
// @Router /estudiante-cursos/test [get]
func (c *EstudianteCursoController) Testing(ctx *fiber.Ctx) error {
	fmt.Println("Endpoint de prueba para estudiante-cursos")
	return ctx.SendString("Hola, este es un endpoint de prueba para estudiante-cursos")
}

// Inscribir godoc
// @Summary Inscribir estudiante en curso
// @Description Inscribe a un estudiante en un curso específico
// @Tags estudiante-cursos
// @Accept json
// @Produce json
// @Param inscripcion body InscripcionRequest true "Datos de inscripción"
// @Success 201 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /estudiante-cursos [post]
func (c *EstudianteCursoController) Inscribir(ctx *fiber.Ctx) error {
	var req InscripcionRequest
	body := ctx.Body()
	fmt.Println("Raw request body:", string(body))

	// Parsear el cuerpo de la solicitud
	if err := ctx.BodyParser(&req); err != nil {
		fmt.Println("Error parsing request body:", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Llamar al servicio para inscribir al estudiante
	err := c.service.Inscribir(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Responder con éxito
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Estudiante inscrito correctamente"})
}

// Desinscribir godoc
// @Summary Desinscribir estudiante de curso
// @Description Elimina la inscripción de un estudiante en un curso específico
// @Tags estudiante-cursos
// @Accept json
// @Produce json
// @Param desinscripcion body DesinscripcionRequest true "Datos de desinscripción"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /estudiante-cursos [delete]
func (c *EstudianteCursoController) Desinscribir(ctx *fiber.Ctx) error {
	var req DesinscripcionRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Llamar al servicio para desinscribir al estudiante
	err := c.service.Desinscribir(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Responder con éxito
	return ctx.JSON(fiber.Map{"message": "Estudiante desinscrito correctamente"})
}

// GetCursosByEstudiante godoc
// @Summary Obtener cursos de un estudiante
// @Description Obtiene todos los cursos en los que está inscrito un estudiante específico
// @Tags estudiante-cursos
// @Produce json
// @Param id path int true "ID del estudiante"
// @Success 200 {array} EstudianteCurso
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /estudiante-cursos/estudiante/{id} [get]
func (c *EstudianteCursoController) GetCursosByEstudiante(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid student ID"})
	}

	// Obtener los cursos en los que está inscrito el estudiante
	cursos, err := c.service.GetCursosDelEstudiante(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Responder con la lista de cursos
	return ctx.JSON(cursos)
}

// GetEstudiantesByCurso godoc
// @Summary Obtener estudiantes de un curso
// @Description Obtiene todos los estudiantes inscritos en un curso específico
// @Tags estudiante-cursos
// @Produce json
// @Param codigo path string true "Código del curso"
// @Success 200 {array} EstudianteCurso
// @Failure 500 {object} fiber.Map
// @Router /estudiante-cursos/curso/{codigo} [get]
func (c *EstudianteCursoController) GetEstudiantesByCurso(ctx *fiber.Ctx) error {
	codigo := ctx.Params("codigo")
	estudiantes, err := c.service.GetEstudiantesByCurso(ctx.Context(), codigo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Responder con la lista de estudiantes
	return ctx.JSON(estudiantes)
}
