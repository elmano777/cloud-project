// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: students_courses.sql

package generated

import (
	"context"
	"database/sql"
)

const createInscripcion = `-- name: CreateInscripcion :execresult
INSERT INTO
    estudiante_cursos (estudiante_id, curso_codigo)
VALUES
    (?, ?)
`

type CreateInscripcionParams struct {
	EstudianteID int32
	CursoCodigo  int32
}

func (q *Queries) CreateInscripcion(ctx context.Context, arg CreateInscripcionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createInscripcion, arg.EstudianteID, arg.CursoCodigo)
}

const deleteInscripcion = `-- name: DeleteInscripcion :exec
DELETE FROM estudiante_cursos
WHERE
    estudiante_id = ?
    AND curso_codigo = ?
`

type DeleteInscripcionParams struct {
	EstudianteID int32
	CursoCodigo  int32
}

func (q *Queries) DeleteInscripcion(ctx context.Context, arg DeleteInscripcionParams) error {
	_, err := q.db.ExecContext(ctx, deleteInscripcion, arg.EstudianteID, arg.CursoCodigo)
	return err
}

const deleteInscripcionesByCurso = `-- name: DeleteInscripcionesByCurso :exec
DELETE FROM estudiante_cursos
WHERE
    curso_codigo = ?
`

func (q *Queries) DeleteInscripcionesByCurso(ctx context.Context, cursoCodigo int32) error {
	_, err := q.db.ExecContext(ctx, deleteInscripcionesByCurso, cursoCodigo)
	return err
}

const deleteInscripcionesByEstudiante = `-- name: DeleteInscripcionesByEstudiante :exec
DELETE FROM estudiante_cursos
WHERE
    estudiante_id = ?
`

func (q *Queries) DeleteInscripcionesByEstudiante(ctx context.Context, estudianteID int32) error {
	_, err := q.db.ExecContext(ctx, deleteInscripcionesByEstudiante, estudianteID)
	return err
}

const getInscripcion = `-- name: GetInscripcion :one
SELECT
    estudiante_id, curso_codigo, inscrito_en
FROM
    estudiante_cursos
WHERE
    estudiante_id = ?
    AND curso_codigo = ?
LIMIT
    1
`

type GetInscripcionParams struct {
	EstudianteID int32
	CursoCodigo  int32
}

func (q *Queries) GetInscripcion(ctx context.Context, arg GetInscripcionParams) (EstudianteCurso, error) {
	row := q.db.QueryRowContext(ctx, getInscripcion, arg.EstudianteID, arg.CursoCodigo)
	var i EstudianteCurso
	err := row.Scan(&i.EstudianteID, &i.CursoCodigo, &i.InscritoEn)
	return i, err
}

const listInscripcionesByCursoWithPagination = `-- name: ListInscripcionesByCursoWithPagination :many
SELECT
    estudiante_id, curso_codigo, inscrito_en
FROM
    estudiante_cursos
WHERE
    curso_codigo = ?
ORDER BY
    inscrito_en
LIMIT
    ?
`

type ListInscripcionesByCursoWithPaginationParams struct {
	CursoCodigo int32
	Limit       int32
}

func (q *Queries) ListInscripcionesByCursoWithPagination(ctx context.Context, arg ListInscripcionesByCursoWithPaginationParams) ([]EstudianteCurso, error) {
	rows, err := q.db.QueryContext(ctx, listInscripcionesByCursoWithPagination, arg.CursoCodigo, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EstudianteCurso
	for rows.Next() {
		var i EstudianteCurso
		if err := rows.Scan(&i.EstudianteID, &i.CursoCodigo, &i.InscritoEn); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listInscripcionesByEstudianteWithPagination = `-- name: ListInscripcionesByEstudianteWithPagination :many
SELECT
    estudiante_id, curso_codigo, inscrito_en
FROM
    estudiante_cursos
WHERE
    estudiante_id = ?
ORDER BY
    inscrito_en
LIMIT
    ?
`

type ListInscripcionesByEstudianteWithPaginationParams struct {
	EstudianteID int32
	Limit        int32
}

func (q *Queries) ListInscripcionesByEstudianteWithPagination(ctx context.Context, arg ListInscripcionesByEstudianteWithPaginationParams) ([]EstudianteCurso, error) {
	rows, err := q.db.QueryContext(ctx, listInscripcionesByEstudianteWithPagination, arg.EstudianteID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []EstudianteCurso
	for rows.Next() {
		var i EstudianteCurso
		if err := rows.Scan(&i.EstudianteID, &i.CursoCodigo, &i.InscritoEn); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
