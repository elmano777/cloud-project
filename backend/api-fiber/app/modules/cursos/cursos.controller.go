package cursos

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CursoController struct {
	cursoService *CursoService
}

func NewCursoController(cursoService *CursoService) *CursoController {
	return &CursoController{cursoService: cursoService}
}

// Testing godoc
// @Summary Endpoint de prueba para cursos
// @Description Verifica si el controlador de cursos está funcionando correctamente
// @Tags cursos
// @Produce plain
// @Success 200 {string} string "CursoController funcionando correctamente"
// @Router /cursos/test [get]
func (c *CursoController) Testing(ctx *fiber.Ctx) error {
	return ctx.SendString("CursoController funcionando correctamente")
}

// CreateCurso godoc
// @Summary Crear un nuevo curso
// @Description Crea un nuevo curso en el sistema
// @Tags cursos
// @Accept json
// @Produce json
// @Param curso body CreateCursoRequest true "Datos del curso"
// @Success 201 {object} CursoResponse
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /cursos [post]
func (c *CursoController) CreateCurso(ctx *fiber.Ctx) error {
	var req CreateCursoRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Petición inválida"})
	}

	curso, err := c.cursoService.CreateCurso(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(curso)
}

// GetCursoByCodigo godoc
// @Summary Obtener un curso por su código
// @Description Obtiene la información de un curso específico según su código
// @Tags cursos
// @Produce json
// @Param codigo path string true "Código del curso"
// @Success 200 {object} CursoResponse
// @Failure 400 {object} fiber.Map
// @Failure 404 {object} fiber.Map
// @Router /cursos/{codigo} [get]
func (c *CursoController) GetCursoByCodigo(ctx *fiber.Ctx) error {
	codigoStr := ctx.Params("codigo")

	// Convertir el código de string a int
	codigo, err := strconv.Atoi(codigoStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Código de curso inválido"})
	}

	curso, err := c.cursoService.GetCursoByCodigo(ctx.Context(), codigo)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Curso no encontrado"})
	}

	return ctx.JSON(curso)
}

// ListCursos godoc
// @Summary Listar todos los cursos
// @Description Obtiene una lista paginada de todos los cursos disponibles
// @Tags cursos
// @Produce json
// @Success 200 {array} CursoResponse
// @Failure 500 {object} fiber.Map
// @Router /cursos [get]
func (c *CursoController) ListCursos(ctx *fiber.Ctx) error {
	result, err := c.cursoService.ListCursos(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(result)
}

// UpdateCurso godoc
// @Summary Actualizar un curso existente
// @Description Actualiza la información de un curso según su código
// @Tags cursos
// @Accept json
// @Produce json
// @Param codigo path string true "Código del curso"
// @Param curso body UpdateCursoRequest true "Datos actualizados del curso"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /cursos/{codigo} [put]
func (c *CursoController) UpdateCurso(ctx *fiber.Ctx) error {
	codigoStr := ctx.Params("codigo")

	// Convertir el código de string a int
	codigo, err := strconv.Atoi(codigoStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Código de curso inválido"})
	}

	var req UpdateCursoRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Petición inválida"})
	}

	if err := c.cursoService.UpdateCurso(ctx.Context(), codigo, req); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Curso actualizado correctamente"})
}

// DeleteCurso godoc
// @Summary Eliminar un curso
// @Description Elimina un curso según su código
// @Tags cursos
// @Produce json
// @Param codigo path string true "Código del curso"
// @Success 200 {object} fiber.Map
// @Failure 400 {object} fiber.Map
// @Failure 500 {object} fiber.Map
// @Router /cursos/{codigo} [delete]
func (c *CursoController) DeleteCurso(ctx *fiber.Ctx) error {
	codigoStr := ctx.Params("codigo")

	// Convertir el código de string a int
	codigo, err := strconv.Atoi(codigoStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Código de curso inválido"})
	}

	if err := c.cursoService.DeleteCurso(ctx.Context(), codigo); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Curso eliminado exitosamente"})
}
