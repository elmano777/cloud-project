package estudiantecurso

import "context"

type EstudianteCursoService struct {
	repo       *EstudianteCursoRepository
	userClient *UserClient
}

func NewEstudianteCursoService(repo *EstudianteCursoRepository, userClient *UserClient) *EstudianteCursoService {
	return &EstudianteCursoService{repo: repo, userClient: userClient}
}

func (s *EstudianteCursoService) Inscribir(ctx context.Context, req InscripcionRequest) error {
	if err := s.userClient.ValidarUsuario(req.EstudianteID); err != nil {
		return err
	}

	return s.repo.InscribirEstudiante(ctx, req.EstudianteID, req.CursoCodigo)
}

func (s *EstudianteCursoService) Desinscribir(ctx context.Context, req DesinscripcionRequest) error {
	return s.repo.DeleteInscripcion(ctx, req.EstudianteID, req.CursoCodigo)
}

func (s *EstudianteCursoService) GetCursosDelEstudiante(ctx context.Context, estudianteID, limit, page int) ([]EstudianteCurso, error) {
	generatedCursos, err := s.repo.GetCursosByEstudiante(ctx, estudianteID, limit, page)
	if err != nil {
		return nil, err
	}

	// Convert from generated.EstudianteCurso to EstudianteCurso
	cursos := make([]EstudianteCurso, len(generatedCursos))
	for i, curso := range generatedCursos {
		cursos[i] = EstudianteCurso{
			CursoCodigo:  int(curso.CursoCodigo),
			EstudianteID: int(curso.EstudianteID),
			InscritoEn:   curso.InscritoEn.Time, // Asegúrate que el tipo de generated.EstudianteCurso.InscritoEn sea compatible
			// con tu EstudianteCurso.InscritoEn. sqlc usualmente usa time.Time o sql.NullTime.
		}
	}

	return cursos, nil
}

func (s *EstudianteCursoService) GetEstudiantesByCurso(ctx context.Context, cursoCodigo int, limit, page int) ([]EstudianteCurso, error) {
	generatedEstudiantes, err := s.repo.GetEstudiantesByCurso(ctx, cursoCodigo, limit, page)
	if err != nil {
		return nil, err
	}

	// Convert from generated.EstudianteCurso to EstudianteCurso
	estudiantes := make([]EstudianteCurso, len(generatedEstudiantes))
	for i, estudiante := range generatedEstudiantes {
		estudiantes[i] = EstudianteCurso{
			CursoCodigo:  int(estudiante.CursoCodigo),
			EstudianteID: int(estudiante.EstudianteID),
			InscritoEn:   estudiante.InscritoEn.Time, // Misma consideración que arriba para el tipo.
		}
	}

	return estudiantes, nil
}
