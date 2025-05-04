package cursos

import (
	"api-fiber/app/utils"
)

// CreateCursoRequest - DTO para creación de cursos
type CreateCursoRequest struct {
	Codigo  int    `json:"codigo" validate:"required,min=5,max=20"`
	Nombre  string `json:"nombre" validate:"required,min=3,max=100"`
	Horario string `json:"horario" validate:"required,min=5,max=50"`
	Ciclo   string `json:"ciclo" validate:"required,min=5,max=20"` // Ejemplo: "2023-1", "2023-2"
}

// UpdateCursoRequest - DTO para actualización de cursos
type UpdateCursoRequest struct {
	Codigo  *int    `json:"codigo,omitempty" validate:"omitempty,min=5,max=20"`
	Nombre  *string `json:"nombre,omitempty" validate:"omitempty,min=3,max=100"`
	Horario *string `json:"horario,omitempty" validate:"omitempty,min=5,max=50"`
	Ciclo   *string `json:"ciclo,omitempty" validate:"omitempty,min=5,max=20"`
}

// CursoResponse - DTO para respuesta de curso
type CursoResponse struct {
	Codigo           int              `json:"codigo"`
	Nombre           string           `json:"nombre"`
	Horario          string           `json:"horario"`
	Ciclo            string           `json:"ciclo"`
	CreatedAt        utils.CustomDate `json:"created_at"`
	EstudiantesCount int              `json:"estudiantes_count,omitempty"`
}

// ListCursosResponse - DTO para listado de cursos
type ListCursosResponse struct {
	Data  []CursoResponse `json:"data"`
	Total int             `json:"total"`
	Page  int             `json:"page"`
	Limit int             `json:"limit"`
}
