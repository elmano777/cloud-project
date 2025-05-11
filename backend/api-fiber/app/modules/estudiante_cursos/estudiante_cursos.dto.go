package estudiantecurso

type InscripcionRequest struct {
	EstudianteID int `json:"estudiante_id" validate:"required"`
	CursoCodigo  int `json:"curso_codigo" validate:"required,min=5,max=20"`
}

type CursoResponse struct {
	Codigo  int    `json:"codigo"`
	Nombre  string `json:"nombre"`
	Horario string `json:"horario"`
	Ciclo   string `json:"ciclo"`
}

// ListCursosResponse - DTO para la respuesta del listado de cursos
type ListCursosResponse struct {
	Data  []CursoResponse `json:"data"`
	Total int             `json:"total"`
	Page  int             `json:"page"`
	Limit int             `json:"limit"`
}

type DesinscripcionRequest struct {
	EstudianteID int `json:"estudiante_id" validate:"required"`
	CursoCodigo  int `json:"curso_codigo" validate:"required,min=5,max=20"`
}
