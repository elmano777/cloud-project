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

// Testing endpoint para verificar si el controller está funcionando
func (c *EstudianteCursoController) Testing(ctx *fiber.Ctx) error {
	fmt.Println("Endpoint de prueba para estudiante-cursos")
	return ctx.SendString("Hola, este es un endpoint de prueba para estudiante-cursos")
}

// Inscribir un estudiante en un curso
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

// Eliminar inscripción de un estudiante en un curso
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

// Obtener cursos de un estudiante por ID
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

// Obtener estudiantes inscritos en un curso
func (c *EstudianteCursoController) GetEstudiantesByCurso(ctx *fiber.Ctx) error {
	codigo := ctx.Params("codigo")
	estudiantes, err := c.service.GetEstudiantesByCurso(ctx.Context(), codigo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	// Responder con la lista de estudiantes
	return ctx.JSON(estudiantes)
}
