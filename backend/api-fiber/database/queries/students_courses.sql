-- name: GetInscripcion :one
SELECT * FROM estudiante_cursos
WHERE estudiante_id = ? AND curso_codigo = ? LIMIT 1;

-- name: ListInscripcionesByEstudiante :many
SELECT * FROM estudiante_cursos
WHERE estudiante_id = ?
ORDER BY inscrito_en;

-- name: ListInscripcionesByCurso :many
SELECT * FROM estudiante_cursos
WHERE curso_codigo = ?
ORDER BY inscrito_en;

-- name: CreateInscripcion :execresult
INSERT INTO estudiante_cursos (
  estudiante_id, curso_codigo
) VALUES (
  ?, ?
);

-- name: DeleteInscripcion :exec
DELETE FROM estudiante_cursos
WHERE estudiante_id = ? AND curso_codigo = ?;

-- name: DeleteInscripcionesByEstudiante :exec
DELETE FROM estudiante_cursos
WHERE estudiante_id = ?;

-- name: DeleteInscripcionesByCurso :exec
DELETE FROM estudiante_cursos
WHERE curso_codigo = ?;