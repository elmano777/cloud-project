package cursos

import (
	"context"

	"api-fiber/database/generated"
)

type CursoRepository struct {
	queries *generated.Queries
}

func NewCursoRepository(queries *generated.Queries) *CursoRepository {
	return &CursoRepository{queries: queries}
}

func (r *CursoRepository) GetCursoByCodigo(ctx context.Context, codigo int) (*generated.Curso, error) {
	curso, err := r.queries.GetCurso(ctx, int32(codigo))
	if err != nil {
		return nil, err
	}
	return &curso, nil
}

func (r *CursoRepository) CreateCurso(ctx context.Context, params generated.CreateCursoParams) (*generated.Curso, error) {
	_, err := r.queries.CreateCurso(ctx, params)
	if err != nil {
		return nil, err
	}

	curso, err := r.queries.GetCurso(ctx, params.Codigo)
	if err != nil {
		return nil, err
	}
	return &curso, nil
}

func (r *CursoRepository) ListCursos(ctx context.Context, limit, offset int) ([]generated.Curso, error) {
	return r.queries.ListCursosWithPagination(ctx, generated.ListCursosWithPaginationParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
}

func (r *CursoRepository) UpdateCurso(ctx context.Context, params generated.UpdateCursoParams) error {
	return r.queries.UpdateCurso(ctx, params)
}

func (r *CursoRepository) DeleteCurso(ctx context.Context, codigo int) error {
	return r.queries.DeleteCurso(ctx, int32(codigo))
}
