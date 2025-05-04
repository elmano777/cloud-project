-- name: GetCurso :one
SELECT
    *
FROM
    cursos
WHERE
    codigo = ?
LIMIT
    1;

-- name: ListCursos :many
SELECT
    *
FROM
    cursos
ORDER BY
    nombre;

-- name: CreateCurso :execresult
INSERT INTO
    cursos (codigo, nombre, horario, ciclo)
VALUES
    (?, ?, ?, ?);

-- name: UpdateCurso :exec
UPDATE cursos
SET
    nombre = ?,
    horario = ?,
    ciclo = ?
WHERE
    codigo = ?;

-- name: DeleteCurso :exec
DELETE FROM cursos
WHERE
    codigo = ?;
