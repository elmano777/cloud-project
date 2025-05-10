package docs

type ErrorResponse struct {
	Error string `json:"error" example:"Mensaje de error"`
}

type SuccessResponse struct {
	Message string `json:"message" example:"Operación realizada con éxito"`
}

type SwaggerCursoRequest struct {
	Codigo  int    `json:"codigo" example:"10001" format:"int"`
	Nombre  string `json:"nombre" example:"Programación en Go"`
	Horario string `json:"horario" example:"Lunes y Miércoles 14:00-16:00"`
	Ciclo   string `json:"ciclo" example:"2023-2"`
}

type SwaggerCursoResponse struct {
	Codigo    int    `json:"codigo" example:"10001" format:"int"`
	Nombre    string `json:"nombre" example:"Programación en Go"`
	Horario   string `json:"horario" example:"Lunes y Miércoles 14:00-16:00"`
	Ciclo     string `json:"ciclo" example:"2023-2"`
	CreatedAt string `json:"created_at" example:"2023-09-01" format:"date"`
}

type SwaggerCursoUpdateRequest struct {
	Nombre  string `json:"nombre,omitempty" example:"Programación Avanzada en Go"`
	Horario string `json:"horario,omitempty" example:"Lunes y Miércoles 16:00-18:00"`
	Ciclo   string `json:"ciclo,omitempty" example:"2023-2"`
}

type SwaggerInscripcionRequest struct {
	EstudianteID int    `json:"estudiante_id" example:"12345" format:"int"`
	CursoCodigo  string `json:"curso_codigo" example:"10001"`
}

type SwaggerEstudianteCursoResponse struct {
	EstudianteID int    `json:"estudiante_id" example:"12345" format:"int"`
	CursoCodigo  string `json:"curso_codigo" example:"10001"`
	InscritoEn   string `json:"inscrito_en" example:"2023-09-05T10:30:00Z" format:"date-time"`
}
