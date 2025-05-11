-- name: GetInscripcion :one
SELECT
    *
FROM
    estudiante_cursos
WHERE
    estudiante_id = ?
    AND curso_codigo = ?
LIMIT
    1;

-- name: ListInscripcionesByEstudianteWithPagination :many
SELECT
    *
FROM
    estudiante_cursos
WHERE
    estudiante_id = ?
ORDER BY
    inscrito_en
LIMIT
    ?
OFFSET
    ?;

-- name: ListInscripcionesByCursoWithPagination :many
SELECT
    *
FROM
    estudiante_cursos
WHERE
    curso_codigo = ?
ORDER BY
    inscrito_en
LIMIT
    ?
OFFSET
    ?;

-- name: CreateInscripcion :execresult
INSERT INTO
    estudiante_cursos (estudiante_id, curso_codigo)
VALUES
    (?, ?);

-- name: DeleteInscripcion :exec
DELETE FROM estudiante_cursos
WHERE
    estudiante_id = ?
    AND curso_codigo = ?;

-- name: DeleteInscripcionesByEstudiante :exec
DELETE FROM estudiante_cursos
WHERE
    estudiante_id = ?;

-- name: DeleteInscripcionesByCurso :exec
DELETE FROM estudiante_cursos
WHERE
    curso_codigo = ?;
