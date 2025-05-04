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

// Endpoint de prueba
func (c *CursoController) Testing(ctx *fiber.Ctx) error {
	return ctx.SendString("CursoController funcionando correctamente")
}

// Crear un nuevo curso
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

// Obtener un curso por su código
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

// Listar cursos paginados
func (c *CursoController) ListCursos(ctx *fiber.Ctx) error {
	result, err := c.cursoService.ListCursos(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(result)
}

// Actualizar un curso existente
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

// Eliminar un curso
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
