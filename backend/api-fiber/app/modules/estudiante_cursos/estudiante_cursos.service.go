package estudiantecurso

import "context"

type EstudianteCursoService struct {
	repo       *EstudianteCursoRepository
	userClient *UserClient
}

func NewEstudianteCursoService(repo *EstudianteCursoRepository, userClient *UserClient) *EstudianteCursoService {
	return &EstudianteCursoService{repo: repo, userClient: userClient}
}

// Inscribir - Inscribe a un estudiante en un curso
func (s *EstudianteCursoService) Inscribir(ctx context.Context, req InscripcionRequest) error {
	// Validar si el estudiante existe en el microservicio de usuarios
	if err := s.userClient.ValidarUsuario(req.EstudianteID); err != nil {
		return err
	}

	// Inscribir al estudiante en el curso
	return s.repo.InscribirEstudiante(ctx, req.EstudianteID, req.CursoCodigo)
}

// Desinscribir - Elimina la inscripción de un estudiante de un curso
func (s *EstudianteCursoService) Desinscribir(ctx context.Context, req DesinscripcionRequest) error {
	return s.repo.DeleteInscripcion(ctx, req.EstudianteID, req.CursoCodigo)
}

// GetCursosDelEstudiante - Devuelve los cursos a los que un estudiante está inscrito
// GetCursosDelEstudiante - Devuelve los cursos a los que un estudiante está inscrito
func (s *EstudianteCursoService) GetCursosDelEstudiante(ctx context.Context, estudianteID int) ([]EstudianteCurso, error) {
	cursos, err := s.repo.GetCursosByEstudiante(ctx, estudianteID)
	if err != nil {
		return nil, err
	}

	// Convertir de []generated.EstudianteCurso a []EstudianteCurso
	resultado := make([]EstudianteCurso, len(cursos))
	for i, curso := range cursos {
		resultado[i] = EstudianteCurso{
			EstudianteID: int(curso.EstudianteID),
			CursoCodigo:  curso.CursoCodigo,
		}
	}

	return resultado, nil
}

// GetEstudiantesByCurso - Devuelve los estudiantes inscritos en un curso específico
func (s *EstudianteCursoService) GetEstudiantesByCurso(ctx context.Context, cursoCodigo string) ([]EstudianteCurso, error) {
	estudiantes, err := s.repo.GetEstudiantesByCurso(ctx, cursoCodigo)
	if err != nil {
		return nil, err
	}

	// Convertir de []generated.EstudianteCurso a []EstudianteCurso
	resultado := make([]EstudianteCurso, len(estudiantes))
	for i, estudiante := range estudiantes {
		resultado[i] = EstudianteCurso{
			EstudianteID: int(estudiante.EstudianteID),
			CursoCodigo:  estudiante.CursoCodigo,
		}
	}

	return resultado, nil
}
