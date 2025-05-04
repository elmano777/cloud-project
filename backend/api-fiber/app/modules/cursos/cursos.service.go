package cursos

import (
	"api-fiber/app/utils"
	"api-fiber/database/generated"
	"context"
	"errors"
	"time"
)

type CursoService struct {
	repo *CursoRepository
}

func NewCursoService(repo *CursoRepository) *CursoService {
	return &CursoService{repo: repo}
}

func (s *CursoService) CreateCurso(ctx context.Context, req CreateCursoRequest) (*CursoResponse, error) {
	params := generated.CreateCursoParams{
		Codigo:  int32(req.Codigo),
		Nombre:  req.Nombre,
		Horario: req.Horario,
		Ciclo:   req.Ciclo,
	}

	curso, err := s.repo.CreateCurso(ctx, params)
	if err != nil {
		return nil, err
	}

	createdAt := utils.CustomDate(time.Time{}) // default zero date
	if curso.CreatedAt.Valid {
		createdAt = utils.CustomDate(curso.CreatedAt.Time)
	}

	return &CursoResponse{
		Codigo:    int(curso.Codigo),
		Nombre:    curso.Nombre,
		Horario:   curso.Horario,
		Ciclo:     curso.Ciclo,
		CreatedAt: createdAt,
	}, nil

}

func (s *CursoService) GetCursoByCodigo(ctx context.Context, codigo int) (*CursoResponse, error) {
	curso, err := s.repo.GetCursoByCodigo(ctx, codigo)
	if err != nil {
		return nil, err
	}

	createdAt := utils.CustomDate(time.Time{})
	if curso.CreatedAt.Valid {
		createdAt = utils.CustomDate(curso.CreatedAt.Time)
	}

	return &CursoResponse{
		Codigo:    int(curso.Codigo),
		Nombre:    curso.Nombre,
		Horario:   curso.Horario,
		Ciclo:     curso.Ciclo,
		CreatedAt: createdAt,
	}, nil
}

func (s *CursoService) ListCursos(ctx context.Context) ([]CursoResponse, error) {
	cursos, err := s.repo.ListCursos(ctx)
	if err != nil {
		return nil, err
	}

	var response []CursoResponse

	for _, curso := range cursos {
		createdAt := utils.CustomDate(time.Time{})
		if curso.CreatedAt.Valid {
			createdAt = utils.CustomDate(curso.CreatedAt.Time)
		}

		response = append(response, CursoResponse{
			Codigo:    int(curso.Codigo),
			Nombre:    curso.Nombre,
			Horario:   curso.Horario,
			Ciclo:     curso.Ciclo,
			CreatedAt: createdAt,
		})
	}

	return response, nil
}

func (s *CursoService) UpdateCurso(ctx context.Context, codigo int, req UpdateCursoRequest) error {
	// Validamos que el curso exista
	_, err := s.repo.GetCursoByCodigo(ctx, codigo)
	if err != nil {
		return errors.New("curso no encontrado")
	}

	params := generated.UpdateCursoParams{
		Codigo: int32(codigo),
	}

	if req.Nombre != nil {
		params.Nombre = *req.Nombre
	}
	if req.Horario != nil {
		params.Horario = *req.Horario
	}
	if req.Ciclo != nil {
		params.Ciclo = *req.Ciclo
	}

	return s.repo.UpdateCurso(ctx, params)
}

func (s *CursoService) DeleteCurso(ctx context.Context, codigo int) error {
	return s.repo.DeleteCurso(ctx, codigo)
}
