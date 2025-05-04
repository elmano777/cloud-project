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
func (r *EstudianteCursoRepository) InscribirEstudiante(ctx context.Context, estudianteID int, cursoCodigo string) error {
	_, err := r.queries.CreateInscripcion(ctx, generated.CreateInscripcionParams{
		EstudianteID: int32(estudianteID),
		CursoCodigo:  cursoCodigo,
	})
	return err
}

// GetCursosByEstudiante - Devuelve los cursos a los que un estudiante está inscrito
func (r *EstudianteCursoRepository) GetCursosByEstudiante(ctx context.Context, estudianteID int) ([]generated.EstudianteCurso, error) {
	return r.queries.ListInscripcionesByEstudiante(ctx, int32(estudianteID))
}

// DeleteInscripcion - Elimina la inscripción de un estudiante en un curso
func (r *EstudianteCursoRepository) DeleteInscripcion(ctx context.Context, estudianteID int, cursoCodigo string) error {
	return r.queries.DeleteInscripcion(ctx, generated.DeleteInscripcionParams{
		EstudianteID: int32(estudianteID),
		CursoCodigo:  cursoCodigo,
	})
}

// GetEstudiantesByCurso - Devuelve los estudiantes inscritos en un curso específico
func (r *EstudianteCursoRepository) GetEstudiantesByCurso(ctx context.Context, cursoCodigo string) ([]generated.EstudianteCurso, error) {
	return r.queries.ListInscripcionesByCurso(ctx, cursoCodigo)
}

// DeleteInscripcionesByCurso - Elimina todas las inscripciones de un curso
func (r *EstudianteCursoRepository) DeleteInscripcionesByCurso(ctx context.Context, cursoCodigo string) error {
	return r.queries.DeleteInscripcionesByCurso(ctx, cursoCodigo)
}

// DeleteInscripcionesByEstudiante - Elimina todas las inscripciones de un estudiante
func (r *EstudianteCursoRepository) DeleteInscripcionesByEstudiante(ctx context.Context, estudianteID int) error {
	return r.queries.DeleteInscripcionesByEstudiante(ctx, int32(estudianteID))
}

// GetInscripcion - Obtiene una inscripción específica (estudiante, curso)
func (r *EstudianteCursoRepository) GetInscripcion(ctx context.Context, estudianteID int, cursoCodigo string) (generated.EstudianteCurso, error) {
	return r.queries.GetInscripcion(ctx, generated.GetInscripcionParams{
		EstudianteID: int32(estudianteID),
		CursoCodigo:  cursoCodigo,
	})
}
