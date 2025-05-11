package estudiantecurso

import (
	"api-fiber/database/generated"
	"context"
)

type EstudianteCursoRepository struct {
	queries *generated.Queries
}

func NewEstudianteCursoRepository(queries *generated.Queries) *EstudianteCursoRepository {
	return &EstudianteCursoRepository{queries: queries}
}

// InscribirEstudiante - Inscribe a un estudiante en un curso
func (r *EstudianteCursoRepository) InscribirEstudiante(ctx context.Context, estudianteID int, cursoCodigo int) error {
	_, err := r.queries.CreateInscripcion(ctx, generated.CreateInscripcionParams{
		EstudianteID: int32(estudianteID),
		CursoCodigo:  int32(cursoCodigo),
	})
	return err
}

// GetCursosByEstudiante - Devuelve los cursos a los que un estudiante está inscrito
func (r *EstudianteCursoRepository) GetCursosByEstudiante(ctx context.Context, estudianteID, limit, offset int) ([]generated.EstudianteCurso, error) {
	return r.queries.ListInscripcionesByEstudianteWithPagination(ctx, generated.ListInscripcionesByEstudianteWithPaginationParams{
		EstudianteID: int32(estudianteID),
		Limit:        int32(limit),
		Offset:       int32(offset),
	})
}

// DeleteInscripcion - Elimina la inscripción de un estudiante en un curso
func (r *EstudianteCursoRepository) DeleteInscripcion(ctx context.Context, estudianteID int, cursoCodigo int) error {
	return r.queries.DeleteInscripcion(ctx, generated.DeleteInscripcionParams{
		EstudianteID: int32(estudianteID),
		CursoCodigo:  int32(cursoCodigo),
	})
}

// GetEstudiantesByCurso - Devuelve los estudiantes inscritos en un curso específico
func (r *EstudianteCursoRepository) GetEstudiantesByCurso(ctx context.Context, cursoCodigo int, limit, offset int) ([]generated.EstudianteCurso, error) {
	return r.queries.ListInscripcionesByCursoWithPagination(ctx, generated.ListInscripcionesByCursoWithPaginationParams{
		CursoCodigo: int32(cursoCodigo),
		Limit:       int32(limit),
		Offset:      int32(offset),
	})
}

// DeleteInscripcionesByCurso - Elimina todas las inscripciones de un curso
func (r *EstudianteCursoRepository) DeleteInscripcionesByCurso(ctx context.Context, cursoCodigo int) error {
	return r.queries.DeleteInscripcionesByCurso(ctx, int32(cursoCodigo))
}

// DeleteInscripcionesByEstudiante - Elimina todas las inscripciones de un estudiante
func (r *EstudianteCursoRepository) DeleteInscripcionesByEstudiante(ctx context.Context, estudianteID int) error {
	return r.queries.DeleteInscripcionesByEstudiante(ctx, int32(estudianteID))
}

// GetInscripcion - Obtiene una inscripción específica (estudiante, curso)
func (r *EstudianteCursoRepository) GetInscripcion(ctx context.Context, estudianteID int, cursoCodigo int) (generated.EstudianteCurso, error) {
	return r.queries.GetInscripcion(ctx, generated.GetInscripcionParams{
		EstudianteID: int32(estudianteID),
		CursoCodigo:  int32(cursoCodigo),
	})
}
